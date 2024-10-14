package main

import (
	"context"
	"time"
)

func GatherData(ctx context.Context, args Args) ([]file, error) {
    if ctx.Err() != nil {
        return nil, err
    }
 
    localCtx, localCancel := context.WithTimeout(ctx, 2 * time.Second)
    local, err := getFilesLocal(localCtx, args.local)
    localCancel()
    if err != nil {
        return nil, err
    }

    remoteCtx, remoteCancel := context.WithTimeout(ctx, 3 *time.Second)
    remote, err := getFilesRemote(remoteCtx, args.remote)
    remoteCancel()
    if err != nil {
        return nil, err
    }
    return append(local, remote), nil
} 