/*
 * Copyright (c) 2021-present Fabien Potencier <fabien@symfony.com>
 *
 * This file is part of Symfony CLI project
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package process

import (
	"net"
	"os"
	"syscall"
)

// FindAvailablePort finds an available port
func FindAvailablePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	// in case of ENOTAVAIL, fallback to try to bind on 127.0.0.1 only
	if err != nil && ErrIsENOTAVAIL(err) {
		l, err = net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: int(0)})
	}
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func ErrIsENOTAVAIL(err error) bool {
	if op, ok := err.(*net.OpError); ok {
		err = op.Err
	}
	if sys, ok := err.(*os.SyscallError); ok {
		err = sys.Err
	}
	return err == syscall.EADDRNOTAVAIL
}
