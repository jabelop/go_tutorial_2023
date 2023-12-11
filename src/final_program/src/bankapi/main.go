package main

import (
	"bank/bankcore"
	"fmt"
	"net/http"
	"strconv"
	"log"

)

var accounts = map[float64]*bankcore.Account{}

func statement(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            fmt.Fprintf(w, account.Statement())
        }
    }
}

func deposit(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            err := account.Deposit(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

func withdraw(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            err := account.Withdraw(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

func transfer(w http.ResponseWriter, req *http.Request) {
    origin := req.URL.Query().Get("origin")
    destination := req.URL.Query().Get("destination")
    amountqs := req.URL.Query().Get("amount")

    if origin == "" {
        fmt.Fprintf(w, "Account origin is missing!")
        return
    }

    if destination == "" {
        fmt.Fprintf(w, "Account destination is missing!")
        return
    }

    if originNumber , err := strconv.ParseFloat(origin, 64); err != nil {
        fmt.Fprintf(w, "Invalid origin account number! %v", originNumber)
    } else if destinationNumber, err := strconv.ParseFloat(destination, 64); err != nil {
        fmt.Fprintf(w, "Invalid destination account number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {         
        originAccount, okOrigin := accounts[originNumber]
        destinationAccount, okDestination := accounts[destinationNumber]

        if !okOrigin {
            fmt.Fprintf(w, "Account with number %v can't be found!", originNumber)
        } else if !okDestination {
            fmt.Fprintf(w, "Account with number %v can't be found!", destinationNumber)
        } else {
            errTransfer := originAccount.Transfer(destinationAccount, amount)
            if errTransfer != nil {
                fmt.Fprintf(w, "Error: $%v", errTransfer)
            } else {
                fmt.Fprintf(w, "%v => %s | %v => %s",originNumber, originAccount.Statement(), destinationNumber, destinationAccount.Statement())
            }
        }
    }
}

func main() {
    accounts[1001] = &bankcore.Account{
        Customer: bankcore.Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number: 1001,
    }

    accounts[1002] = &bankcore.Account{
        Customer: bankcore.Customer{
            Name:    "Marie",
            Address: "San Diego, California",
            Phone:   "(216) 557 1133",
        },
        Number: 1002,
    }

	http.HandleFunc("/statement", statement)
    http.HandleFunc("/deposit", deposit)
    http.HandleFunc("/withdraw", withdraw)
    http.HandleFunc("/transfer", transfer)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
