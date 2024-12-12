package main

import (
	"context"
	"fmt"
	"github.com/ilkamo/jupiter-go/jupiter"
	"strconv"
	"time"
)

func main() {
	jupClient, err := jupiter.NewClientWithResponses(jupiter.DefaultAPIURL)
	if err != nil {
		panic(err)
	}
	base := "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
	inputAmount := int64(10000000)
	routers := make([]string, 0)
	routers = append(routers, "So11111111111111111111111111111111111111112")
	for true {
		for _, router := range routers {
			err := arb(jupClient, base, router, inputAmount)
			if err != nil {
				fmt.Printf("%s \n", err)
			}
			time.Sleep(time.Second * 1)
		}
	}
}

func arb(jupClient *jupiter.ClientWithResponses, base string, router string, inputAmount int64) error {
	inputAmount1 := inputAmount
	quoteResponse1, err := jupClient.GetQuoteWithResponse(context.Background(), &jupiter.GetQuoteParams{
		InputMint:                     base,
		OutputMint:                    router,
		Amount:                        int(inputAmount1),
		SlippageBps:                   nil,
		AutoSlippage:                  nil,
		AutoSlippageCollisionUsdValue: nil,
		ComputeAutoSlippage:           nil,
		MaxAutoSlippageBps:            nil,
		SwapMode:                      nil,
		Dexes:                         nil,
		ExcludeDexes:                  nil,
		RestrictIntermediateTokens:    nil,
		OnlyDirectRoutes:              nil,
		AsLegacyTransaction:           nil,
		PlatformFeeBps:                nil,
		MaxAccounts:                   nil,
		MinimizeSlippage:              nil,
		PreferLiquidDexes:             nil,
	})
	if err != nil {
		return err
	}
	if quoteResponse1 == nil || quoteResponse1.JSON200 == nil {
		return fmt.Errorf("empty response")
	}

	outputAmountStr1 := quoteResponse1.JSON200.OutAmount
	outputAmount1, err := strconv.ParseInt(outputAmountStr1, 10, 64)
	if err != nil {
		return err
	}

	quoteResponse2, err := jupClient.GetQuoteWithResponse(context.Background(), &jupiter.GetQuoteParams{
		InputMint:                     router,
		OutputMint:                    base,
		Amount:                        int(outputAmount1),
		SlippageBps:                   nil,
		AutoSlippage:                  nil,
		AutoSlippageCollisionUsdValue: nil,
		ComputeAutoSlippage:           nil,
		MaxAutoSlippageBps:            nil,
		SwapMode:                      nil,
		Dexes:                         nil,
		ExcludeDexes:                  nil,
		RestrictIntermediateTokens:    nil,
		OnlyDirectRoutes:              nil,
		AsLegacyTransaction:           nil,
		PlatformFeeBps:                nil,
		MaxAccounts:                   nil,
		MinimizeSlippage:              nil,
		PreferLiquidDexes:             nil,
	})
	if err != nil {
		return err
	}
	if quoteResponse2 == nil || quoteResponse2.JSON200 == nil {
		return fmt.Errorf("empty response")
	}
	outputAmountStr2 := quoteResponse2.JSON200.OutAmount
	outputAmount2, err := strconv.ParseInt(outputAmountStr2, 10, 64)
	if err != nil {
		return err
	}

	fmt.Printf("swap %t, input amount: %d output amount: %d\n", outputAmount2 > inputAmount1, inputAmount1, outputAmount2)
	return nil
}
