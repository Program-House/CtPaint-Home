module Main.View exposing (view)

import Html exposing (..)
import Html.Attributes exposing (..)
import Main.Model exposing (Model)
import Main.Message exposing (Msg(..))
import View.Components exposing (words)
import View.AppBar as AppBar
import SignIn.View as SignIn


-- VIEW


view : Model -> Html Msg
view model =
    if model.loggedIn then
        div
            [ class "main logged-in" ]
            [ AppBar.view model ]
    else
        div
            [ class "main" ]
            [ SignIn.view model ]