use std::{cell::RefCell, collections::HashMap, rc::Rc};




pub enum Event {
    Nothing,
    Invalid,
    Win,
}


#[derive(Clone)]
pub struct Tile {
    //pub update: bool,
    pub color: bool,
    pub group: Rc<Box<u8>>,
}


pub struct Board {
    //pub groups: HashMap<u8, Rc<Box<u8>>>,
    pub tiles: [[Option<Tile>; 6]; 7],
}
impl Board {
    pub fn new() -> Self {
        const INIT: Option<Tile> = None;
        const INIT_ONE: [Option<Tile>; 6] = [INIT; 6];
        Board {
            tiles: [INIT_ONE; 7],
        }
    }


    fn update_groups(&mut self, color: bool, cur: (usize, usize), left: (usize, usize), right: (usize, usize)) -> Event {

        Event::Nothing

    }


    pub fn place(&mut self, color: bool, column: usize) -> Event {

        let mut found = None;

        // find next open tile in this column
        for row in 0..6 {
            let cur_col = &mut self.tiles[column as usize];
            if let Some(_) = cur_col[row] { //ignore all filled spaces
                continue;
            }

            //let tile = &mut cur_col[row].unwrap();

            // have an open space
            found = Some(row);
        }

        let row;
        match found {
            None => return Event::Invalid,
            Some(i) => row = i,
        }

        //let mut search = [true; 6];

        //if column == 0 { //on left side
        //    search[0] = false;
        //    search[1] = false;
        //    search[2] = false;
        //} else if column == 6 { //on right side
        //    search[3] = false;
        //    search[4] = false;
        //    search[5] = false;
        //}

        //if row == 0 { //on bottom
        //    search[2] = false;
        //    search[5] = false;
        //} else if row == 5 { //on top
        //    search[0] = false;
        //    search[3] = false;
        //}

        //let left_tile = &mut self.tiles[column+1][row-1];
        //let right_tile = &mut self.tiles[column-1][row+1];

        // ul and dr
        //let code: u8;
        //if !search[0] { //left side is dead, only go right side
        //    code = match &mut self.tiles[column+1][row-1] {
        //        Some(tile) => if tile.color == color { 2 } else { 0 },
        //        None => 0,
        //    };

        //} else if !search[5] { //right side is dead, only go left side
        //    code = match &mut self.tiles[column-1][row+1] {
        //        Some(tile) => if tile.color == color { 1 } else { 0 },
        //        None => 0,
        //    };

        //} else { //both sides are open
        //    let left_code = match &mut self.tiles[column-1][row+1] {
        //        Some(tile) => if tile.color == color { 1 } else { 0 },
        //        None => 0,
        //    };
        //    let right_code = match &mut self.tiles[column+1][row-1] {
        //        Some(tile) => if tile.color == color { 2 } else { 0 },
        //        None => 0,
        //    };
        //    code = left_code + right_code;
        //}

        for ((l_col, l_row), (r_col, r_row)) in [((-1, 1), (1, -1)), ((-1, 0), (1, 0)), ((-1, -1), (1, 1))] {

            // first figure out where we can check without out-of-bounds errors
            let mut search = [true; 6];

            if column == 0 { //on left side
                search[0] = false;
                search[1] = false;
                search[2] = false;
            } else if column == 6 { //on right side
                search[3] = false;
                search[4] = false;
                search[5] = false;
            }

            if row == 0 { //on bottom
                search[2] = false;
                search[5] = false;
            } else if row == 5 { //on top
                search[0] = false;
                search[3] = false;
            }

            let left_col = (column as i32 + l_col) as usize;
            let left_row = (row as i32 + l_row) as usize;
            let right_col = (column as i32 + r_col) as usize;
            let right_row = (row as i32 + r_row) as usize;

            let code: u8;
            if !search[0] { //left side is dead, only go right side
                code = match &mut self.tiles[right_col][right_row] {
                    Some(tile) => if tile.color == color { 2 } else { 0 },
                    None => 0,
                };

            } else if !search[5] { //right side is dead, only go left side
                code = match &mut self.tiles[left_col][left_row] {
                    Some(tile) => if tile.color == color { 1 } else { 0 },
                    None => 0,
                };

            } else { //both sides are open
                let left_code = match &mut self.tiles[left_col][left_row] {
                    Some(tile) => if tile.color == color { 1 } else { 0 },
                    None => 0,
                };
                let right_code = match &mut self.tiles[right_col][right_row] {
                    Some(tile) => if tile.color == color { 2 } else { 0 },
                    None => 0,
                };
                code = left_code + right_code;
            }

            // handle codes 0, 1, 2, 3
            match code {
                0 => { //create new group, join neither

                },
                1 => { //join left

                },
                2 => { //join right

                },
                3 => { //unify both sides

                },
                _ => unreachable!(),
            }
            
        }

        // handle codes 0, 1, 2, 3

        //match code {
        //    0 => {

        //    },
        //    1 => {

        //    },
        //    2 => {

        //    },
        //    3 => {

        //    },
        //    _ => unreachable!(),
        //}


        todo!();


        

        // this column is completely filled up
        Event::Nothing
    }
}


