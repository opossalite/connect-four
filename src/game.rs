use crate::{engine::Board, printing::{print_board, prompt_move}};



pub fn player_player() {
    let mut board = Board::new();
    let mut input = String::new();
    let mut player = true;

    loop {
        print_board(&mut board);
        let result = prompt_move(player);
        print!("{}", result);

        player = !player;
        break;
    }
}





