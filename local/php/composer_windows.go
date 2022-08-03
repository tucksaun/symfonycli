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

package php

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func findComposerSystemSpecific(extraBin string) string {
	// Special Support for Scoop
	scoopPath := os.Getenv("SCOOP")
	if scoopPath == "" {
		if homedir, err := homedir.Dir(); err != nil {
			scoopPath = filepath.Join(homedir, "scoop")
		}
	}

	scoopGlobalPath := os.Getenv("SCOOP_GLOBAL")
	if scoopGlobalPath == "" {
		if programData := os.Getenv("PROGRAMDATA"); programData != "" {
			scoopGlobalPath = filepath.Join(programData, "scoop")
		}
	}

	for _, path := range []string{scoopPath, scoopGlobalPath} {
		if path == "" {
			continue
		}

		pharPath := filepath.Join(path, "apps", "composer", "current", "composer.phar")
		d, err := os.Stat(pharPath)
		if err != nil {
			continue
		}
		if m := d.Mode(); !m.IsDir() {
			// Yep!
			return pharPath
		}
	}

	return ""
}
