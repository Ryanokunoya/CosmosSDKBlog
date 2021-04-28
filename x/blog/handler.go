package blog

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/example/blog/x/blog/keeper"
	"github.com/example/blog/x/blog/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
    // this line is used by starport scaffolding # handler/msgServer

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		// x/blog/handler.go

		switch msg := msg.(type) {
			// this line is used by starport scaffolding # 1
			case *types.MsgCreatePost:
				return handleMsgCreatePost(ctx, k, msg)
			default:
	
	}
}
