package player

import (
    "testing"
)

func TestPlayerSetup(t *testing.T) {
    var player Player

    if player.Name != "" {
        t.Error("player does not initialize with an empty name,",
            "initializes with name:", player.Name)
    }

    if player.Char != 0 {
        t.Error("player does not initialize with an empty char,",
            "initializes with char:", string(player.Char))
    }

    // player.SetupPlayer("Player1", []string{}, []string{})
    //
    // if player.Name != "Test" {
    //     t.Error("player name does not match set name Test,",
    //         "name is:", player.Name)
    // }
    //
    // if player.Char != 'X' {
    //     t.Error("player char does not match set char X,",
    //         "char is:", string(player.Char))
    // }
}

func Test_stringSearch(t *testing.T) {
    strs := []string{
        "test",
        "t",
        "test2",
    }

    if !searchString(strs, "test") {
        t.Error("search cannot find string when it exists in slice")
    }

    if searchString(strs, "nope") {
        t.Error("search returns true when search does not exists in slice")
    }
}
