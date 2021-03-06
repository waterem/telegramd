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
	"fmt"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
)

/*
 rpc_requst:
	body: { photos_getUserPhotos
	  user_id: { inputUserSelf },
	  offset: 1 [INT],
	  max_id: 0 [LONG],
	  limit: 5 [INT],
	},

 rpc_result:
  body: { rpc_result
    req_msg_id: 6537205080566771468 [LONG],
    result: { photos_photosSlice
      count: 1 [INT],
      photos: [ vector<0x0> ],
      users: [ vector<0x0> ],
    },
  },

 */
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (s *PhotosServiceImpl) PhotosGetUserPhotos(ctx context.Context, request *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("PhotosGetUserPhotos - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PhotosGetUserPhotos logic

	return nil, fmt.Errorf("Not impl PhotosGetUserPhotos")
}
