package main

import (
	"backend/models"
	"errors"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"reflect"
	"strconv"
	"fmt"
	"math/rand"

	"github.com/darahayes/go-boom"
	"github.com/patrickmn/go-cache"
	valid "github.com/asaskevich/govalidator"
	"github.com/nu7hatch/gouuid"
)
var cacheManager  = cache.New(5*time.Minute, 10*time.Minute)
const numberCards = 52
var ranks  		  = [13]string{"A","2","3","4","5","6","7","8","9","T","J","Q","K"}
var suits  		  = [4]string{"S", "D", "C", "H"}
var 		 	  decks[numberCards]string

func (app *application) createDeckOfCards(w http.ResponseWriter, r *http.Request) {
	/*
	Start creation ordered cards in deck
	*/
	for i := 0; i < len(decks); i++ {
			decks[i] = ranks[i%13] + suits[i/13] 		
	}
	/*
	End creation ordered cardsin deck
	*/

	/*
	Start set of ordered cards
	*/
	uid, err := uuid.NewV4()
	CreateDeckModel := models.CreateDeck{
		DeckID: uid.String(),
		Shuffled: false,
		Remaining: len(decks),
	}

	//setting caches
	cacheManager = cache.New(5*time.Minute, 10*time.Minute)

	cacheManager.Set("deck_id", CreateDeckModel.DeckID, cache.DefaultExpiration)
	cacheManager.Set("shuffled", CreateDeckModel.Shuffled, cache.DefaultExpiration)
	cacheManager.Set("remaining", CreateDeckModel.Remaining, cache.DefaultExpiration)
	cacheManager.Set("deck", decks, cache.DefaultExpiration)
	/*
	End set of ordered cards
	*/

	/*
	Start responses
	*/
	js, err := json.MarshalIndent(CreateDeckModel, "", "\t")
	if err != nil {
		app.logger.Println(err)
		boom.BadRequest(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	/*
	End responses
	*/
}

func (app *application) createDeckOfCardsByCards(w http.ResponseWriter, r *http.Request) {
	/*
	Start validation for negative response
	*/
	cards := r.URL.Query().Get("cards")
	if cards == "" {
		app.logger.Print(errors.New("Param Cards not found!"))
		err := errors.New("Param cards not found or no values in it! Please, try again!")
		boom.BadRequest(w, err)
		return
	}
	/*
	End validation for negative response
	*/

	/*
	Start creation ordered cards in deck
	*/
	for i := 0; i < len(decks); i++ {
			decks[i] = ranks[i%13] + suits[i/13] 		
	}
	/*
	End creation ordered cards in deck
	*/

	validCards := strings.Split(cards, ",")
	validDecks := make([]string, len(validCards))

	//check if cards sent by the user are valid!
	for i:=0; i < len(validCards); i++ {
		for j:=0; j < len(decks); j++{
			if validCards[i] == decks[j]{
				validDecks[i] = decks[j]
			}
		}
	}

	//removing empty string from the array!!
	var finaldecks []string
	for _, str := range validDecks {
		if str != "" {
			finaldecks = append(finaldecks, str)
		}
	}

	/*
	Start validation for negative response
	*/
	if len(finaldecks) == 0 {
		app.logger.Print(errors.New("Invalid values in param!"))
		err := errors.New("Invalid values in param!! Please, try again with valid values!")
		boom.BadRequest(w, err)
		return
	}
	/*
	End validation for negative response
	*/
	
	/*
	Start set of cards
	*/
	uid, err := uuid.NewV4()
	CreateDeckModel := models.CreateDeck{
		DeckID: uid.String(),
		Shuffled: false,
		Remaining: len(finaldecks),
	}

	cacheManager = cache.New(5*time.Minute, 10*time.Minute)

	cacheManager.Set("deck_id", CreateDeckModel.DeckID, cache.DefaultExpiration)
	cacheManager.Set("shuffled", CreateDeckModel.Shuffled, cache.DefaultExpiration)
	cacheManager.Set("remaining", CreateDeckModel.Remaining, cache.DefaultExpiration)
	cacheManager.Set("deck", finaldecks, cache.DefaultExpiration)
	/*
	Start set of ordered cards
	*/

	/*
	Start responses
	*/
	js, err := json.MarshalIndent(CreateDeckModel, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	/*
	End responses
	*/
}

func (app *application) shuffleDeck(w http.ResponseWriter, r *http.Request){
	/*
	Start validation for negative response
	*/
	deck, found := cacheManager.Get("deck")
	if !found {
		app.logger.Print(errors.New("Cache expired or null!"))
		err := errors.New("Cache expired or null!")
		boom.BadRequest(w, err)
		return
	}
	/*
	End validation for negative response
	*/

	/*
	Start Shuffle of decks
	*/
	deckShuffled := make([]string, reflect.ValueOf(deck).Len())
	for i := 0; i < reflect.ValueOf(deck).Len(); i++ {
        j := rand.Intn(i + 1)
		deckShuffled[i], deckShuffled[j] = fmt.Sprintf("%v", reflect.ValueOf(deck).Index(j).Interface()), fmt.Sprintf("%v", reflect.ValueOf(deck).Index(i).Interface()) 
    }
	/*
	End Shuffle of decks
	*/

	/*
	Start sets of shuffled decks
	*/
	shuffleModel := models.Shuffle{
		Shuffled: true,
	}

	//setting caches
	cacheManager.Set("deck", deckShuffled, cache.DefaultExpiration)
	cacheManager.Set("shuffled", true, cache.DefaultExpiration)
	/*
	End sets of shuffled decks
	*/

	/*
	Start responses
	*/
	js, err := json.MarshalIndent(shuffleModel, "", "\t")
	if err != nil {
		app.logger.Println(err)
		boom.BadRequest(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	/*
	End responses
	*/
}

func (app *application) openDeckOfCards(w http.ResponseWriter, r *http.Request){
	/*
	Start validation for negative response
	*/
	deck, found := cacheManager.Get("deck")
	if !found {
		app.logger.Print(errors.New("Cache expired or null!"))
		err := errors.New("Cache expired or null!")
		boom.BadRequest(w, err)
		return
	}
	/*
	End validation for negative response
	*/

	/*
	Start Creation of model
	*/
	cardsOpened := make([]models.Cards, reflect.ValueOf(deck).Len())
	for i := 0; i<reflect.ValueOf(deck).Len(); i++ {
		cardsOpened[i] = models.Cards{Code: fmt.Sprintf("%v", reflect.ValueOf(deck).Index(i).Interface())}
	}
	
	deckUid, _ := cacheManager.Get("deck_id")
	shuffled, _ := cacheManager.Get("shuffled")
	openDeckModel := models.OpenDeck{
		DeckID: reflect.ValueOf(deckUid).String(),
		Shuffled: reflect.ValueOf(shuffled).Bool(),
		Remaining: reflect.ValueOf(deck).Len(),
		Cards: cardsOpened,
	}
	/*
	End Creation of model
	*/

	/*
	Start responses
	*/
	js, err := json.MarshalIndent(openDeckModel, "", "\t")
	if err != nil {
		app.logger.Println(err)
		boom.BadRequest(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	/*
	End responses
	*/
}

func (app *application) drawCards(w http.ResponseWriter, r *http.Request){
	/*
	Start validations for negative responses
	*/
	count := r.URL.Query().Get("count")
	if count == "" {
		app.logger.Print(errors.New("Param count not found!"))
		err := errors.New("Param count not found or no values in it! Please, try again!")
		boom.BadRequest(w, err)
		return
	}

	if !valid.IsNumeric(count) {
		app.logger.Print(errors.New("Param count is not numeric!"))
		err := errors.New("Param count is not numeric! Please, try again with a numeric value!")
		boom.BadRequest(w, err)
		return
	}

	paramCount, _:= strconv.Atoi(count)
	cardsModel := make([]models.Cards, paramCount)
	
	deck, found := cacheManager.Get("deck")
	if !found {
		app.logger.Print(errors.New("Cache expired or null!"))
		err := errors.New("Cache expired or null!")
		boom.BadRequest(w, err)
		return
	}
	
	if paramCount > reflect.ValueOf(deck).Len()  {
		app.logger.Print(errors.New("Param count is bigger then remaining!"))
		err := errors.New("Param count is bigger then remaining! Please, try again with a count value lower then remaining!")
		boom.BadRequest(w, err)
		return
	}
	/*
	End validations for negative responses
	*/
	
	
	/*
	Start creation of model
	*/ 
	for i := 0; i<paramCount; i++ {
		if reflect.ValueOf(deck).Index(i).Interface() != "" { 
			cardsModel[i] = models.Cards{Code: fmt.Sprintf("%v", reflect.ValueOf(deck).Index(i).Interface())}
		}
	}
	/*
	End creation of model
	*/ 

	/*
	Start responses
	*/
	js, err := json.MarshalIndent(cardsModel, "", "\t")
	if err != nil {
		app.logger.Println(err)
		boom.BadRequest(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	/*
	End responses
	*/
}