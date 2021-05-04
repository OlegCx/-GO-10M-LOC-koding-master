package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new social channel API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for social channel API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SocialChannelAcceptInvite social channel accept invite API
*/
func (a *Client) SocialChannelAcceptInvite(params *SocialChannelAcceptInviteParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelAcceptInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelAcceptInviteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.acceptInvite",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.acceptInvite",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelAcceptInviteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelAcceptInviteOK), nil

}

/*
SocialChannelAddParticipants social channel add participants API
*/
func (a *Client) SocialChannelAddParticipants(params *SocialChannelAddParticipantsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelAddParticipantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelAddParticipantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.addParticipants",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.addParticipants",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelAddParticipantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelAddParticipantsOK), nil

}

/*
SocialChannelByID social channel by Id API
*/
func (a *Client) SocialChannelByID(params *SocialChannelByIDParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.byId",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.byId",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelByIDOK), nil

}

/*
SocialChannelByName social channel by name API
*/
func (a *Client) SocialChannelByName(params *SocialChannelByNameParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.byName",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.byName",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelByNameOK), nil

}

/*
SocialChannelCreate social channel create API
*/
func (a *Client) SocialChannelCreate(params *SocialChannelCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.create",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelCreateOK), nil

}

/*
SocialChannelCreateChannelWithParticipants social channel create channel with participants API
*/
func (a *Client) SocialChannelCreateChannelWithParticipants(params *SocialChannelCreateChannelWithParticipantsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelCreateChannelWithParticipantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelCreateChannelWithParticipantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.createChannelWithParticipants",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.createChannelWithParticipants",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelCreateChannelWithParticipantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelCreateChannelWithParticipantsOK), nil

}

/*
SocialChannelDelete social channel delete API
*/
func (a *Client) SocialChannelDelete(params *SocialChannelDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.delete",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.delete",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelDeleteOK), nil

}

/*
SocialChannelFetchActivities social channel fetch activities API
*/
func (a *Client) SocialChannelFetchActivities(params *SocialChannelFetchActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchActivitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchActivities",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchActivities",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchActivitiesOK), nil

}

/*
SocialChannelFetchActivityCount Method SocialChannel.fetchActivityCount
*/
func (a *Client) SocialChannelFetchActivityCount(params *SocialChannelFetchActivityCountParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchActivityCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchActivityCountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchActivityCount",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchActivityCount",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchActivityCountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchActivityCountOK), nil

}

/*
SocialChannelFetchChannels social channel fetch channels API
*/
func (a *Client) SocialChannelFetchChannels(params *SocialChannelFetchChannelsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchChannelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchChannelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchChannels",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchChannels",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchChannelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchChannelsOK), nil

}

/*
SocialChannelFetchFollowedChannelCount social channel fetch followed channel count API
*/
func (a *Client) SocialChannelFetchFollowedChannelCount(params *SocialChannelFetchFollowedChannelCountParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchFollowedChannelCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchFollowedChannelCountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchFollowedChannelCount",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchFollowedChannelCount",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchFollowedChannelCountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchFollowedChannelCountOK), nil

}

/*
SocialChannelFetchFollowedChannels social channel fetch followed channels API
*/
func (a *Client) SocialChannelFetchFollowedChannels(params *SocialChannelFetchFollowedChannelsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchFollowedChannelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchFollowedChannelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchFollowedChannels",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchFollowedChannels",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchFollowedChannelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchFollowedChannelsOK), nil

}

/*
SocialChannelFetchParticipants social channel fetch participants API
*/
func (a *Client) SocialChannelFetchParticipants(params *SocialChannelFetchParticipantsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchParticipantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchParticipantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchParticipants",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchParticipants",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchParticipantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchParticipantsOK), nil

}

/*
SocialChannelFetchPinnedMessages social channel fetch pinned messages API
*/
func (a *Client) SocialChannelFetchPinnedMessages(params *SocialChannelFetchPinnedMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchPinnedMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchPinnedMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchPinnedMessages",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchPinnedMessages",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchPinnedMessagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchPinnedMessagesOK), nil

}

/*
SocialChannelFetchPopularPosts social channel fetch popular posts API
*/
func (a *Client) SocialChannelFetchPopularPosts(params *SocialChannelFetchPopularPostsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchPopularPostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchPopularPostsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchPopularPosts",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchPopularPosts",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchPopularPostsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchPopularPostsOK), nil

}

/*
SocialChannelFetchPopularTopics social channel fetch popular topics API
*/
func (a *Client) SocialChannelFetchPopularTopics(params *SocialChannelFetchPopularTopicsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchPopularTopicsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchPopularTopicsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchPopularTopics",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchPopularTopics",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchPopularTopicsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchPopularTopicsOK), nil

}

/*
SocialChannelFetchProfileFeed social channel fetch profile feed API
*/
func (a *Client) SocialChannelFetchProfileFeed(params *SocialChannelFetchProfileFeedParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchProfileFeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchProfileFeedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchProfileFeed",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchProfileFeed",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchProfileFeedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchProfileFeedOK), nil

}

/*
SocialChannelFetchProfileFeedCount social channel fetch profile feed count API
*/
func (a *Client) SocialChannelFetchProfileFeedCount(params *SocialChannelFetchProfileFeedCountParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelFetchProfileFeedCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelFetchProfileFeedCountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.fetchProfileFeedCount",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.fetchProfileFeedCount",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelFetchProfileFeedCountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelFetchProfileFeedCountOK), nil

}

/*
SocialChannelGlancePinnedPost social channel glance pinned post API
*/
func (a *Client) SocialChannelGlancePinnedPost(params *SocialChannelGlancePinnedPostParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelGlancePinnedPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelGlancePinnedPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.glancePinnedPost",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.glancePinnedPost",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelGlancePinnedPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelGlancePinnedPostOK), nil

}

/*
SocialChannelLeave social channel leave API
*/
func (a *Client) SocialChannelLeave(params *SocialChannelLeaveParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelLeaveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelLeaveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.leave",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.leave",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelLeaveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelLeaveOK), nil

}

/*
SocialChannelListParticipants social channel list participants API
*/
func (a *Client) SocialChannelListParticipants(params *SocialChannelListParticipantsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelListParticipantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelListParticipantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.listParticipants",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.listParticipants",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelListParticipantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelListParticipantsOK), nil

}

/*
SocialChannelPinMessage social channel pin message API
*/
func (a *Client) SocialChannelPinMessage(params *SocialChannelPinMessageParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelPinMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelPinMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.pinMessage",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.pinMessage",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelPinMessageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelPinMessageOK), nil

}

/*
SocialChannelRejectInvite social channel reject invite API
*/
func (a *Client) SocialChannelRejectInvite(params *SocialChannelRejectInviteParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelRejectInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelRejectInviteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.rejectInvite",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.rejectInvite",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelRejectInviteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelRejectInviteOK), nil

}

/*
SocialChannelRemoveParticipants social channel remove participants API
*/
func (a *Client) SocialChannelRemoveParticipants(params *SocialChannelRemoveParticipantsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelRemoveParticipantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelRemoveParticipantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.removeParticipants",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.removeParticipants",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelRemoveParticipantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelRemoveParticipantsOK), nil

}

/*
SocialChannelSearchTopics social channel search topics API
*/
func (a *Client) SocialChannelSearchTopics(params *SocialChannelSearchTopicsParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelSearchTopicsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelSearchTopicsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.searchTopics",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.searchTopics",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelSearchTopicsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelSearchTopicsOK), nil

}

/*
SocialChannelUnpinMessage social channel unpin message API
*/
func (a *Client) SocialChannelUnpinMessage(params *SocialChannelUnpinMessageParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelUnpinMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelUnpinMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.unpinMessage",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.unpinMessage",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelUnpinMessageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelUnpinMessageOK), nil

}

/*
SocialChannelUpdate social channel update API
*/
func (a *Client) SocialChannelUpdate(params *SocialChannelUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.update",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.update",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelUpdateOK), nil

}

/*
SocialChannelUpdateLastSeenTime social channel update last seen time API
*/
func (a *Client) SocialChannelUpdateLastSeenTime(params *SocialChannelUpdateLastSeenTimeParams, authInfo runtime.ClientAuthInfoWriter) (*SocialChannelUpdateLastSeenTimeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSocialChannelUpdateLastSeenTimeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SocialChannel.updateLastSeenTime",
		Method:             "POST",
		PathPattern:        "/remote.api/SocialChannel.updateLastSeenTime",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SocialChannelUpdateLastSeenTimeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SocialChannelUpdateLastSeenTimeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}