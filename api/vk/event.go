package vk

// CallbackEvent is base event
type CallbackEvent struct {
	// ID of group this event occured in
	GroupID int
	// Secret for Callback API
	Secret string
	// Event itself
	//
	// One of Confirmation, MessageNew, MessageReply, MessageEdit,
	// MessageAllow, MessageDeny, MessageTypingState, PhotoNew,
	// PhotoCommentNew, PhotoCommentEdit, PhotoCommentRestore,
	// PhotoCommentDelete, AudioNew, VideoNew, VideoCommentNew,
	// VideoCommentEdit, VideoCommentRestore, VideoCommentDelete,
	// WallPostNew, WallRepost, WallReplyNew, WallReplyEdit,
	// WallReplyRestore, WallReplyDelete, BoardPostNew, BoardPostEdit,
	// BoardPostRestore, BoardPostDelete, MarketCommentNew,
	// MarketCommentEdit, MarketCommentRestore, MarketCommentDelete,
	// GroupLeave, GroupJoin, UserBlock, UserUnblock, PollVoteNew,
	// GroupOfficersEdit, GroupChangeSettings, GroupChangePhoto,
	// LeadFormsNew, NewVKPayTransaction.
	Event interface{}
}

// MessageNew -- new message is recieved
//
//easyjson:json
type MessageNew struct {
	Message
}

// Confirmation is used in Callback API.
// It requires listener to reply with Confirmation token instead of normal "ok".
//
//easyjson:json
type Confirmation struct{}
