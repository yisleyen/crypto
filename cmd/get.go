/*
Copyright © 2022 Yusuf işleyen <me@yusufisleyen.com>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yisleyen/crypto/parser"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Bu Cli Tool ile, kripto para birimi bilgileri anlık olarak çekilir",
	Long: `
	Bir tanımlama yapılmaz ise, varsayılan olarak bitcoin'in değeri döner. Eğer tanımlama yapılır ise 
	"ripple" veya "shiba-inu" gibi kripto para birimi bilgisi Doviz.com üzerinden anlık olarak çekilir.
	
	*Doviz.com'da bir sorun olması durumunda, bu servis hizmet veremez.'
	
	Örnek Kullanım:
	.crypto get bitcoin  => Bitcoin  [2022-06-15 22:16:26] : $21.670,45
	.crypto get ethereum => Ethereum [2022-06-15 22:16:26] : $1.173,37
	`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		cryptocurrency := "bitcoin"

		if len(args) > 0 {
			cryptocurrency = args[0]
		}

		cryptocurrency = strings.Title(strings.ToLower(cryptocurrency))

		cryptocurrencys := parser.ParseWebPage(cryptocurrency)

		fmt.Printf("%s: %s$ (%s TRY)", cryptocurrency, cryptocurrencys["dollar"], cryptocurrencys["try"])
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
