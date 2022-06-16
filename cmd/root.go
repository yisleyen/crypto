/*
Copyright © 2022 Yusuf işleyen <me@yusufisleyen.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Bu Cli Tool ile, kripto para birimi bilgileri anlık olarak çekilir",
	Long: `
	Bir tanımlama yapılmaz ise, varsayılan olarak bitcoin'in değeri döner. Eğer tanımlama yapılır ise 
	"ripple" veya "shiba-inu" gibi kripto para birimi bilgisi Doviz.com üzerinden anlık olarak çekilir.
	
	*Doviz.com'da bir sorun olması durumunda, bu servis hizmet veremez.'
	
	Örnek Kullanım:
	.crypto get bitcoin  => Bitcoin: 20841,65$ (366018,0 TRY)
	.crypto get ethereum => Ethereum: 1102,48$ (19345,0 TRY)
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.crypto.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
