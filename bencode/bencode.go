package bencode

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
)

//I will be using github.com/jackpal/bencode-go to parse the .torrent file

type bencodeInfo struct {
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"` //binary blob of hashes of each piece
}

type parsedBencode struct {
	Announce     string      `bencode:"announce"`
	Info         bencodeInfo `bencode:"info"`
	Comment      string      `bencode:"comment"`
	CreationDate int         `bencode:"creation date"`
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func Open(path string) (TorrentFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return TorrentFile{}, err
	}
	defer file.Close()

	torrent := parsedBencode{}
	err = bencode.Unmarshal(file, &torrent)
	if err != nil {
		return TorrentFile{}, err
	}
	return torrent.toTorrentFile()
}

func (i *bencodeInfo) hash() ([20]byte, error) {
	buf := new(bytes.Buffer)
	// var buf bytes.Buffer
	err := bencode.Marshal(buf, *i)
	if err != nil {
		return [20]byte{}, err
	}
	h := sha1.Sum(buf.Bytes())
	return h, nil
}

func (i *bencodeInfo) splitPieceHashes() ([][20]byte, error) {
	hashLen := 20 // Length of SHA-1 hash
	buf := []byte(i.Pieces)

	if len(buf)%hashLen != 0 { //ensures buf has valid data by making sure buf is a multiple of 20
		err := fmt.Errorf("len err %d", len(buf))
		return nil, err
	}

	numHashes := len(buf) / hashLen
	hashes := make([][20]byte, numHashes) //hashes slice of [20]bytes, with len numHashes

	for i := 0; i < numHashes; i++ {
		copy(hashes[i][:], buf[i*hashLen:(i+1)*hashLen]) //copies hashes of individual pieces into newly made slcec of [20]bytes
	}
	return hashes, nil
}

func (pb *parsedBencode) toTorrentFile() (TorrentFile, error) {
	infoHash, err := pb.Info.hash()
	if err != nil {
		return TorrentFile{}, err
	}
	pieceHashes, err := pb.Info.splitPieceHashes()
	if err != nil {
		return TorrentFile{}, err
	}
	t := TorrentFile{
		Announce:    pb.Announce,
		InfoHash:    infoHash,
		PieceHashes: pieceHashes,
		PieceLength: pb.Info.PieceLength,
		Length:      pb.Info.Length,
		Name:        pb.Info.Name,
	}
	return t, nil
}
