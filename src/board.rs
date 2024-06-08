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


    pub fn place(&mut self, color: bool, column: usize) -> Event {

        // find next open tile
        for row in 0..3 {
            let cur_col = &mut self.tiles[column];
            if let Some(_) = cur_col[row] { //ignore all filled spaces
                continue;
            }

            //let tile = &mut cur_col[row].unwrap();

            // have an open space, figure out where we can look
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

            // ul and dr
            if !search[0] { //left side is dead, only go right side
                let right_color = match &mut self.tiles[column+1][row-1] {
                    Some(tile) => if tile.color { 1 } else { 2 },
                    None => 0,
                };
                todo!()

            } else if !search[5] { //right side is dead, only go left side
                let left_color = match &mut self.tiles[column-1][row+1] {
                    Some(tile) => if tile.color { 1 } else { 2 },
                    None => 0,
                };
                todo!()

            } else { //both sides are open
                let left_color = match &mut self.tiles[column-1][row+1] {
                    Some(tile) => if tile.color { 1 } else { 2 },
                    None => 0,
                };
                let right_color = match &mut self.tiles[column+1][row-1] {
                    Some(tile) => if tile.color { 1 } else { 2 },
                    None => 0,
                };
                todo!()
            }


        }

        // this column is completely filled up
        Event::Invalid
    }
}


