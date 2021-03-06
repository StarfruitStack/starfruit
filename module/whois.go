/*
 * Copyright 2014 The starfruit Authors. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package module

import (
	"github.com/flatpeach/starfruit/message"
	"github.com/flatpeach/starfruit/server"
	"github.com/flatpeach/starfruit/user"
	"strings"
)

type Whois struct{}

func (module *Whois) Handle(s *server.Server, u *user.User, m *message.Message) error {
	// WHOIS <mask> *( "," <mask> )

	if len(m.Params) != 1 {
		u.SendMessage(message.New(
			s.Config.Server.Name,
			message.ERR_NEEDMOREPARAMS,
			nil,
			"Need more params",
		))

		return nil
	}

	nicks := strings.Split(m.Params[0], ",")
	for _, nick := range nicks {
		target := s.GetUserByNickName(nick)
		if u == nil {
			// @Todo: fulfill the errors here
			continue
		}

		u.SendMessage(message.New(
			s.Config.Server.Name,
			message.RPL_WHOISUSER,
			[]string{
				u.NickName,
				target.NickName,
				target.UserName,
				target.HostName,
				"*",
			},
			u.RealName,
		))

		u.SendMessage(message.New(
			s.Config.Server.Name,
			message.RPL_WHOISSERVER,
			[]string{
				u.NickName,
				target.NickName,
				s.Config.Server.Name,
			},
			s.Config.Server.Name,
		))

		joinedChannels := s.GetJoinedChannels(u.Id)
		if len(joinedChannels) > 0 {
			u.SendMessage(message.New(
				s.Config.Server.Name,
				message.RPL_WHOISCHANNELS,
				[]string{
					u.NickName,
					target.NickName,
				},
				strings.Join((func() []string {
					var names []string
					for _, cnl := range joinedChannels {
						names = append(names, cnl.String())
					}
					return names
				})(), " "),
			))
		}

		u.SendMessage(message.New(
			s.Config.Server.Name,
			message.RPL_ENDOFWHOIS,
			[]string{
				u.NickName,
				target.NickName,
			},
			"End of /WHOIS list.",
		))
	}

	return nil
}
