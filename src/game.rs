use crate::{engine::Board, printing::{print_board, prompt_move}};



pub fn player_player() {
    let mut board = Board::new();
    let mut input = String::new();
    let mut player = true;

    loop {
        print_board(&mut board);
        //let column_raw = prompt_move(player);
        //let column = column_raw.parse::<usize>().unwrap(); //safe unwrap
        let column = prompt_move(player);
        let result = board.place(player, column); //do action on board
        println!("{:?}", result);

        player = !player;
        break;
    }
}





