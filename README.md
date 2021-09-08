﻿# bittorent-client-go

I will be following the blog post [Building a BitTorrent client from the ground up in Go](https://blog.jse.li/posts/torrent/) which I found on [build-your-own-x](https://github.com/danistefanovic/build-your-own-x)

I'll be documenting my progress on [this notion page](https://paurana.notion.site/Bittorent-Client-abd95ac4b113485facd300a3cbd63ed1)

At the moment, the client only supports those .torrent files whose announce urls follow the http protocol. I plan to come back to this project some day and add support for udp protocol too.

**Clone the repository**

```git clone https://github.com/tech-yush/bittorent-client-go.git```

Download the latest [Debian](https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/#indexlist).torrent file or any other .torrent file whose announce url follows the http protocol. 

```go run main.go debian-11.0.0-amd64-netinst.iso.torrent debian.iso```
