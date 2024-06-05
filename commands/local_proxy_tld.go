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

package commands

import (
	"fmt"
	"regexp"

	"github.com/symfony-cli/console"
	"github.com/symfony-cli/symfony-cli/local/proxy"
	"github.com/symfony-cli/symfony-cli/util"
	"github.com/symfony-cli/terminal"
)

var localProxyTLD = &console.Command{
	Category: "local",
	Name:     "proxy:tld",
	Aliases:  []*console.Alias{{Name: "proxy:tld"}, {Name: "proxy:change:tld"}},
	Usage:    "Display or change the TLD for the proxy",
	Args: []*console.Arg{
		{Name: "tld", Description: "The TLD for the project proxy", Optional: true},
	},
	Action: func(c *console.Context) error {
		homeDir := util.GetHomeDir()
		config, err := proxy.Load(homeDir)
		if err != nil {
			return err
		}

		if c.Args().Present() {
			config.TLD = c.Args().Get("tld")
			if !regexp.MustCompile(`^[a-z]{1,63}$`).MatchString(config.TLD) {
				return fmt.Errorf("the TLD must only contain lowercase letters")
			}
			if err = config.Save(); err != nil {
				return err
			}
		}

		terminal.Printfln("<info>The proxy is configured with the following TLD: %s</>", config.TLD)
		return nil
	},
}
