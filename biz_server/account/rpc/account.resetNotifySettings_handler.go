/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz/core/account"
	"github.com/nebulaim/telegramd/biz/base"
	"github.com/nebulaim/telegramd/biz_server/sync_client"
)

// account.resetNotifySettings#db7e1747 = Bool;
func (s *AccountServiceImpl) AccountResetNotifySettings(ctx context.Context, request *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountResetNotifySettings - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AccountResetNotifySettings logic
	account.ResetNotifySettings(md.UserId)
	peer := &base.PeerUtil{}
	peer.PeerType = base.PEER_ALL
	update := mtproto.NewTLUpdateNotifySettings()
	update.SetPeer(peer.ToNotifyPeer())
	updateSettings := mtproto.NewTLPeerNotifySettings()
	updateSettings.SetShowPreviews(true)
	updateSettings.SetSilent(false)
	updateSettings.SetMuteUntil(0)
	updateSettings.SetSound("default")
	update.SetNotifySettings(updateSettings.To_PeerNotifySettings())

	sync_client.GetSyncClient().PushToUserMeOneUpdateData(md.AuthId, md.SessionId, md.UserId, update.To_Update())

	glog.Infof("AccountResetNotifySettings - reply: {true}")
	return mtproto.ToBool(true), nil
}
