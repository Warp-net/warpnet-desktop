//go:build darwin
// +build darwin

/* License generated by licensor(https://github.com/Marvin9/licensor).

Warpnet - Decentralized Social Network
Copyright (C) 2025 Vadim Filin, https://github.com/Warp-net,
<github.com.mecdy@passmail.net>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

WarpNet is provided “as is” without warranty of any kind, either expressed or implied.
Use at your own risk. The maintainers shall not be liable for any damages or data loss
resulting from the use or misuse of this software.
*/

package embedded

import (
	_ "embed"
	_ "github.com/Warp-net/warpnet-desktop/bin"
)

// This file embeds the whole desktop to a single Golang binary file. DO NOT REMOVE.
//
//go:embed bin/warpnet-desktop-darwin
var desktopBinary []byte

func GetDesktopEmbedded() []byte {
	return desktopBinary
}
