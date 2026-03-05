${{ content_synopsis }} This image will give you a [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) and [distroless](https://github.com/11notes/RTFM/blob/main/linux/container/image/distroless.md) qBittorrent installation for your adventures on the high seas[^1] *arrrr*!

${{ content_arr_stack }}

${{ content_uvp }} Good question! Because ...

${{ github:> [!IMPORTANT] }}
${{ github:> }}* ... this image runs [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) as 1000:1000
${{ github:> }}* ... this image has no shell since it is [distroless](https://github.com/11notes/RTFM/blob/main/linux/container/image/distroless.md)
${{ github:> }}* ... this image is built and compiled from source (no userdocs/qbittorrent-nox!)
${{ github:> }}* ... this image supports 32bit architecture
${{ github:> }}* ... this image is auto updated to the latest version via CI/CD
${{ github:> }}* ... this image has a health check
${{ github:> }}* ... this image runs read-only
${{ github:> }}* ... this image is automatically scanned for CVEs before and after publishing
${{ github:> }}* ... this image is created via a secure and pinned CI/CD process
${{ github:> }}* ... this image is very small
${{ github:> }}* ... this image creates a random password at start if none is set in the config
${{ github:> }}* ... this image support [inline configs](https://github.com/11notes/RTFM/blob/master/linux/container/image/11notes/inline-config.md)

If you value security, simplicity and optimizations to the extreme, then this image might be for you.

${{ content_comparison }}

${{ title_volumes }}
* **${{ json_root }}/etc** - Directory of your qBittorrent.conf and other files
* **${{ json_root }}/var** - Directory of your SQlite database for qBittorrent
* **${{ json_root }}/themes (optional)** - Directory of your alternate themes, VueTorrent is already present

${{ content_compose }}

${{ content_defaults }}
| `login` | admin // qbittorrent | login using default config |
| `AdditionalTrackersURL` | [ngosang/trackerslist](https://raw.githubusercontent.com/ngosang/trackerslist/refs/heads/master/trackers_best.txt) | additional trackers that will be added to every torrent |

${{ content_environment }}
| `QBITTORRENT_USER_AGENT` *(optional)* | sets the user-agent to a custom value if needed |  |
| `QBITTORRENT_PEER_ID` *(optional)* | sets the peer ID to a custom value if needed |  |

${{ content_source }}

${{ content_parent }}

${{ content_built }}

${{ content_tips }}

${{ title_caution }}
${{ github:> [!CAUTION] }}
${{ github:> }}* If you use the image with the default configuration, please make sure to change the default web ui login account password or provide your own qBittorrent.conf!

[^1]: Check the [compose.vpn.yml](https://github.com/11notes/docker-qbittorrent/blob/master/compose.vpn.yml) example on how to use this image with a VPN provider like gluetun.