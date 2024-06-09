use std::io::{self, Write};

use colored::Colorize;

use crate::{engine::{Board, Tile}, game::player_player};



pub fn prompt_menu() {
    std::process::Command::new("clear").status().unwrap();
    println!("What would you like to do?

        a) Train AI
        s) Player vs AI
        d) Player vs Player
        f) Testing / Debugging

        x) Exit\n");
    print!(">>> ");
    io::stdout().flush().unwrap();

    let mut input = String::new();
    std::io::stdin().read_line(&mut input).unwrap();

    match &input.as_str()[0..input.len()-1] {
        "a" => {},
        "s" => {},
        "d" => player_player(),
        "f" => {},
        "x" => {},
        _ => {},
    }
}


pub fn print_board(board: &mut Board) {
    std::process::Command::new("clear").status().unwrap();
    println!();
    println!("  1   2   3   4   5   6   7");
    for j in 0..6 {
        println!("{}", "+---+---+---+---+---+---+---+".blue());
        for i in 0..7 {
            match &mut board.tiles[i][j] {
                //None => print!("{} ⦿ ", "|".blue()),
                None => print!("{}   ", "|".blue()),
                Some(Tile {color: true, groups: _, vertical: _}) => print!("{} {} ", "|".blue(), "⦿".bold().red()),
                Some(Tile {color: false, groups: _, vertical: _}) => print!("{} {} ", "|".blue(), "⦿".bold().yellow()),
            }
        }
        println!("{}", "|".blue());
    }
    println!("{}", "+---+---+---+---+---+---+---+".blue());
    println!("  1   2   3   4   5   6   7\n");
}


pub fn prompt_move(turn: bool) -> usize {
    loop {
        if turn {
            println!("{}'s turn", "red".red());
            print!("{}", ">>> ".red());
        } else {
            println!("{}'s turn", "yellow".yellow());
            print!("{}", ">>> ".yellow());
        }
        io::stdout().flush().unwrap();

        let mut input = String::new();
        std::io::stdin().read_line(&mut input).unwrap();

        match &input.as_str()[0..input.len()-1] {
            "1" => return 0,
            "2" => return 1,
            "3" => return 2,
            "4" => return 3,
            "5" => return 4,
            "6" => return 5,
            "7" => return 6,
            _ => println!("\nBad input, please input a number from 1-7\n"),
        }
    }
}


