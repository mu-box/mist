# Contributor: Hennik Hunsaker <hennikhunsaker@microbox.cloud>
# Maintainer: Hennik Hunsaker <hennikhunsaker@microbox.cloud>
pkgname=mist
pkgver=0.2.3
pkgrel=0
pkgdesc="A distributed, tag-based pub-sub service for modern web applications and container-driven cloud."
url="https://github.com/mu-box/mist"
arch="all"
license="MIT"
depends=""
makedepends="go git bash"
checkdepends=""
install=""
subpackages=""
source=""
srcdir="/tmp/abuild/mist"
builddir=""

build() {
	go get -t -v .
	go install github.com/mitchellh/gox@latest
	export PATH="$(go env | grep GOPATH | sed -E 's/GOPATH="(.*)"/\1/')/bin:${PATH}"
	./scripts/build.sh
}

check() {
	# Replace with proper check command(s)
	:
}

package() {
	install -m 0755 -D ./build/mist "$pkgdir"/sbin/mist
}
