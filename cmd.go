package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chat-app",
	Short: "Serving chat functionality that is backed by ejabberd XMPP server",
}

func Execute() error {
	return rootCmd.Execute()
}
