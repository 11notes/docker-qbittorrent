![banner](https://raw.githubusercontent.com/11notes/static/refs/heads/main/img/banner/README.png)

# QBITTORRENT
![size](https://img.shields.io/badge/image_size-27MB-green?color=%2338ad2d)![5px](https://raw.githubusercontent.com/11notes/static/refs/heads/main/img/markdown/transparent5x2px.png)![pulls](https://img.shields.io/docker/pulls/11notes/qbittorrent?color=2b75d6)![5px](https://raw.githubusercontent.com/11notes/static/refs/heads/main/img/markdown/transparent5x2px.png)[<img src="https://img.shields.io/github/issues/11notes/docker-qbittorrent?color=7842f5">](https://github.com/11notes/docker-qbittorrent/issues)![5px](https://raw.githubusercontent.com/11notes/static/refs/heads/main/img/markdown/transparent5x2px.png)![swiss_made](https://img.shields.io/badge/Swiss_Made-FFFFFF?labelColor=FF0000&logo=data:image/svg%2bxml;base64,PHN2ZyB2ZXJzaW9uPSIxIiB3aWR0aD0iNTEyIiBoZWlnaHQ9IjUxMiIgdmlld0JveD0iMCAwIDMyIDMyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxyZWN0IHdpZHRoPSIzMiIgaGVpZ2h0PSIzMiIgZmlsbD0idHJhbnNwYXJlbnQiLz4KICA8cGF0aCBkPSJtMTMgNmg2djdoN3Y2aC03djdoLTZ2LTdoLTd2LTZoN3oiIGZpbGw9IiNmZmYiLz4KPC9zdmc+)

run qBittorrent rootless and distroless

# INTRODUCTION 📢

[qBittorrent](https://github.com/qbittorrent/qBittorrent) (created by [qbittorrent](https://github.com/qbittorrent)) is a bittorrent client programmed in C++ / Qt that uses libtorrent (sometimes called libtorrent-rasterbar) by Arvid Norberg.

# SYNOPSIS 📖
**What can I do with this?** This image will give you a [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) and [distroless](https://github.com/11notes/RTFM/blob/main/linux/container/image/distroless.md) qBittorrent installation for your adventures on the high seas[^1] *arrrr*!

# ARR STACK IMAGES 🏴‍☠️
This image is part of the so called arr-stack (apps to pirate and manage media content). Here is the list of all it's companion apps for the best pirate experience:

- [11notes/configarr](https://github.com/11notes/docker-configarr) - as your TRaSH guide syncer for Sonarr and Radarr
- [11notes/plex](https://github.com/11notes/docker-plex) - as your media server
- [11notes/prowlarr](https://github.com/11notes/docker-prowlarr) - to manage all your indexers
- [11notes/radarr](https://github.com/11notes/docker-radarr) - to manage your films
- [11notes/sabnzbd](https://github.com/11notes/docker-sabnzbd) - as your usenet client
- [11notes/sonarr](https://github.com/11notes/docker-sonarr) - to manage your TV shows

# UNIQUE VALUE PROPOSITION 💶
**Why should I run this image and not the other image(s) that already exist?** Good question! Because ...

> [!IMPORTANT]
>* ... this image runs [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) as 1000:1000
>* ... this image has no shell since it is [distroless](https://github.com/11notes/RTFM/blob/main/linux/container/image/distroless.md)
>* ... this image is built and compiled from source (no userdocs/qbittorrent-nox!)
>* ... this image supports 32bit architecture
>* ... this image is auto updated to the latest version via CI/CD
>* ... this image has a health check
>* ... this image runs read-only
>* ... this image is automatically scanned for CVEs before and after publishing
>* ... this image is created via a secure and pinned CI/CD process
>* ... this image is very small

If you value security, simplicity and optimizations to the extreme, then this image might be for you.

# COMPARISON 🏁
Below you find a comparison between this image and the most used or original one.

| **image** | **size on disk** | **init default as** | **[distroless](https://github.com/11notes/RTFM/blob/main/linux/container/image/distroless.md)** | supported architectures
| ---: | ---: | :---: | :---: | :---: |
| 11notes/qbittorrent | 27MB | 1000:1000 | ✅ | amd64, arm64, armv7 |
| home-operations/qbittorrent | 111MB | 65534:65533 | ❌ | amd64, arm64 |
| hotio/qbittorrent | 159MB | 0:0 | ❌ | amd64, arm64 |
| qbittorrentofficial/qbittorrent-nox | 167MB | 0:0 | ❌ | 386, amd64, arm64, armv6, armv7, riscv64 |
| linuxserver/qbittorrent | 503MB | 0:0 | ❌ | amd64, arm64 |

# VOLUMES 📁
* **/qbittorrent/etc** - Directory of your qBittorrent.conf and other files
* **/qbittorrent/var** - Directory of your SQlite database for qBittorrent
* **/qbittorrent/themes (optional)** - Directory of your alternate themes, VueTorrent is already present

# COMPOSE ✂️
```yaml
name: "arr"

x-lockdown: &lockdown
  # prevents write access to the image itself
  read_only: true
  # prevents any process within the container to gain more privileges
  security_opt:
    - "no-new-privileges=true"

services:
  qbittorrent:
    image: "11notes/qbittorrent:5.1.4"
    <<: *lockdown
    environment:
      TZ: "Europe/Zurich"
      QBITTORRENT_USER_AGENT: "11notes/qbittorrent" # optional
      QBITTORRENT_PEER_ID: "-qB11-" # optional
    volumes:
      - "qbittorrent.etc:/qbittorrent/etc"
      - "qbittorrent.var:/qbittorrent/var"
    ports:
      - "3000:8080/tcp"
      - "6881:6881/tcp"
      - "6881:6881/udp"
    networks:
      frontend:
    restart: "always"

volumes:
  qbittorrent.etc:
  qbittorrent.var:

networks:
  frontend:
```
To find out how you can change the default UID/GID of this container image, consult the [RTFM](https://github.com/11notes/RTFM/blob/main/linux/container/image/11notes/how-to.changeUIDGID.md#change-uidgid-the-correct-way).

# DEFAULT SETTINGS 🗃️
| Parameter | Value | Description |
| --- | --- | --- |
| `user` | docker | user name |
| `uid` | 1000 | [user identifier](https://en.wikipedia.org/wiki/User_identifier) |
| `gid` | 1000 | [group identifier](https://en.wikipedia.org/wiki/Group_identifier) |
| `home` | /qbittorrent | home directory of user docker |
| `login` | admin // qbittorrent | login using default config |
| `AdditionalTrackersURL` | [ngosang/trackerslist](https://raw.githubusercontent.com/ngosang/trackerslist/refs/heads/master/trackers_best.txt) | additional trackers that will be added to every torrent |

# ENVIRONMENT 📝
| Parameter | Value | Default |
| --- | --- | --- |
| `TZ` | [Time Zone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) | |
| `DEBUG` | Will activate debug option for container image and app (if available) | |
| `QBITTORRENT_USER_AGENT` *(optional)* | sets the user-agent to a custom value if needed |  |
| `QBITTORRENT_PEER_ID` *(optional)* | sets the peer ID to a custom value if needed |  |

# MAIN TAGS 🏷️
These are the main tags for the image. There is also a tag for each commit and its shorthand sha256 value.

* [5.1.4](https://hub.docker.com/r/11notes/qbittorrent/tags?name=5.1.4)
* [5.1.4-unraid](https://hub.docker.com/r/11notes/qbittorrent/tags?name=5.1.4-unraid)

### There is no latest tag, what am I supposed to do about updates?
It is my opinion that the ```:latest``` tag is a bad habbit and should not be used at all. Many developers introduce **breaking changes** in new releases. This would messed up everything for people who use ```:latest```. If you don’t want to change the tag to the latest [semver](https://semver.org/), simply use the short versions of [semver](https://semver.org/). Instead of using ```:5.1.4``` you can use ```:5``` or ```:5.1```. Since on each new version these tags are updated to the latest version of the software, using them is identical to using ```:latest``` but at least fixed to a major or minor version. Which in theory should not introduce breaking changes.

If you still insist on having the bleeding edge release of this app, simply use the ```:rolling``` tag, but be warned! You will get the latest version of the app instantly, regardless of breaking changes or security issues or what so ever. You do this at your own risk!

# REGISTRIES ☁️
```
docker pull 11notes/qbittorrent:5.1.4
docker pull ghcr.io/11notes/qbittorrent:5.1.4
docker pull quay.io/11notes/qbittorrent:5.1.4
```

# UNRAID VERSION 🟠
This image supports unraid by default. Simply add **-unraid** to any tag and the image will run as 99:100 instead of 1000:1000 causing no issues on unraid. Enjoy.

# SOURCE 💾
* [11notes/qbittorrent](https://github.com/11notes/docker-qbittorrent)

# PARENT IMAGE 🏛️
> [!IMPORTANT]
>This image is not based on another image but uses [scratch](https://hub.docker.com/_/scratch) as the starting layer.
>The image consists of the following distroless layers that were added:
>* [11notes/distroless](https://github.com/11notes/docker-distroless/blob/master/arch.dockerfile) - contains users, timezones and Root CA certificates, nothing else
>* [11notes/distroless:localhealth](https://github.com/11notes/docker-distroless/blob/master/localhealth.dockerfile) - app to execute HTTP requests only on 127.0.0.1

# BUILT WITH 🧰
* [qBittorrent](https://github.com/qbittorrent/qBittorrent)

# GENERAL TIPS 📌
> [!TIP]
>* Use a reverse proxy like Traefik, Nginx, HAproxy to terminate TLS and to protect your endpoints
>* Use Let’s Encrypt DNS-01 challenge to obtain valid SSL certificates for your services

# CAUTION ⚠️
> [!CAUTION]
>* If you use the image with the default configuration, please make sure to change the default web ui login account password or provide your own qBittorrent.conf!

[^1]: Check the [compose.vpn.yml](https://github.com/11notes/docker-qbittorrent/blob/master/compose.vpn.yml) example on how to use this image with a VPN provider like gluetun.

# ElevenNotes™️
This image is provided to you at your own risk. Always make backups before updating an image to a different version. Check the [releases](https://github.com/11notes/docker-qbittorrent/releases) for breaking changes. If you have any problems with using this image simply raise an [issue](https://github.com/11notes/docker-qbittorrent/issues), thanks. If you have a question or inputs please create a new [discussion](https://github.com/11notes/docker-qbittorrent/discussions) instead of an issue. You can find all my other repositories on [github](https://github.com/11notes?tab=repositories).

*created 28.11.2025, 11:48:54 (CET)*