use std::{cell::RefCell, mem::MaybeUninit, rc::Rc};



/// Represents the result of placing a tile on the board
pub enum Event {
    Nothing,
    Invalid,
    Win,
}


/// Struct for one tile, of which a board contains 42
#[derive(Clone)]
pub struct Tile {
    pub color: bool,
    pub groups: [Rc<RefCell<u8>>; 3],
    pub vertical: i32,
}


/// Main struct used for Connect Four
pub struct Board {
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


    /// Drop a piece into the board and return the result
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
            None => return Event::Invalid, //this column is completely filled up
            Some(i) => row = i,
        }

        // first handle vertical wins
        let vertical_count;
        if row > 0 {
            vertical_count = self.tiles[column][row-1].as_ref().unwrap().vertical + 1;

            // handle vertical win
            if vertical_count > 3 {
                return Event::Win;
            }

        } else { //first one in the column
            vertical_count = 1;
        }

        // prepare for group management
        const COORDS: [((i32, i32), (i32, i32)); 3] = [((-1, 1), (1, -1)), ((-1, 0), (1, 0)), ((-1, -1), (1, 1))];
        let mut groups: [Rc<RefCell<u8>>; 3] = unsafe { [MaybeUninit::uninit().assume_init(), MaybeUninit::uninit().assume_init(), MaybeUninit::uninit().assume_init()] };
        let mut win = false;

        // iterate through all directions
        for i in 0..COORDS.len() {
            let ((l_col, l_row), (r_col, r_row)) = COORDS[i];

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

            // figure out which code we are dealing with, which update case
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
                    groups[i] = Rc::new(RefCell::new(1)); //set initial size of group to 1
                },
                1 => { //join left
                    let left_group = self.tiles[left_col][left_row].as_ref().unwrap().groups[i].clone();
                    *left_group.borrow_mut() += 1; //increment group size
                    groups[i] = left_group; //add this tile to the group
                },
                2 => { //join right
                    let right_group = self.tiles[right_col][right_row].as_ref().unwrap().groups[i].clone();
                    *right_group.borrow_mut() += 1; //increment group size
                    groups[i] = right_group; //add this tile to the group
                },
                3 => { //unify both sides
                    let left_group = self.tiles[left_col][left_row].as_ref().unwrap().groups[i].clone();
                    let right_group_size = *self.tiles[right_col][right_row].as_ref().unwrap().groups[i].borrow_mut();
                    let summed_size = *left_group.borrow_mut() + right_group_size + 1; //figure out new size
                    *left_group.borrow_mut() = summed_size; //increment group size

                    // update both ends of the former right group to be in the new unified left group
                    let far_right_col = (r_col * right_group_size as i32 + column as i32) as usize;
                    let far_right_row = (r_row * right_group_size as i32 + row as i32) as usize;
                    self.tiles[right_col][right_row].as_mut().unwrap().groups[i] = left_group.clone(); //update close end
                    self.tiles[far_right_col][far_right_row].as_mut().unwrap().groups[i] = left_group.clone(); //update far end

                    // add this tile to the group
                    groups[i] = left_group;

                    // check for win
                    if summed_size > 3 {
                        win = true;
                    }

                },
                _ => unreachable!(),
            }

            // handle win condition
            if win {
                return Event::Win;
            }
        }

        // now that updates have been handled in all directions, create the new tile
        let tile = Tile {
            color,
            groups: [groups[0].clone(), groups[1].clone(), groups[2].clone()],
            vertical: vertical_count,
        };
        self.tiles[column][row] = Some(tile);

        Event::Nothing
    }
}


