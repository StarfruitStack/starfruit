/*
 * Copyright 2014 The starfruit Authors. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package module

import (
	"fmt"
	"github.com/flatpeach/starfruit/message"
	"github.com/flatpeach/starfruit/server"
	"github.com/flatpeach/starfruit/user"
	"time"
)

type Time struct{}

func (module *Time) Handle(s *server.Server, u *user.User, m *message.Message) error {
	// TIME [ <target> ]

	now := time.Now()

	u.SendMessage(message.New(
		s.Config.Server.Name,
		message.RPL_TIME,
		[]string{
			u.NickName,
			s.Config.Server.Name,
		},
		fmt.Sprintf("%s", now.Format("Jan 2, 2006 at 3:04pm (MST)")),
	))

	return nil
}
