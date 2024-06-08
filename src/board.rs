


pub enum Event {
    Nothing,
    Invalid,
    Win,
}


//#[derive(Copy, Clone)]
//pub enum Tile {
//    Empty,
//    ///color,
//    Filled(bool, [u8; 7], [bool; 6]), 
//}


#[derive(Clone, Copy)]
pub struct Tile {
    pub color: bool,
    pub chains: [u8; 7], //ul, l, dl, ur, r, dr, d
    pub update: [bool; 6],
}



pub struct Board {
    pub tiles: [[Option<Tile>; 6]; 7],
}
impl Board {
    pub fn new() -> Self {
        Board {
            tiles: [[None; 6]; 7]
        }
    }


    ///// Attempt to place a piece in the given column
    //pub fn place(&mut self, color: bool, column: usize) -> Event {
    //    match self.tiles[column][5] {
    //        Some(_) => Event::Invalid, //column was filled up, invalid placement
    //        None => self.place_internal(color, column),
    //    }
    //}


    ///// Handle placing a piece in a column (guaranteed to be available)
    /// Attempt to place a piece in the given column
    pub fn place(&mut self, color: bool, column: usize) -> Event {
        let cur_col = &mut self.tiles[column];

        // find next open tile
        for row in 0..6 {
            if let Some(_) = cur_col[row] { //ignore all filled spaces
                continue;
            }

            // open space
            cur_col[row] = Some(Tile {
                color,
                chains: [0, 0, 0, 0, 0, 0, 0],
                update: [true, true, true, true, true, true],
            });
            let tile = &mut cur_col[row].unwrap();

            //// cover some edge cases for enabled-ness
            //if column == 0 { //on left edge
            //    tile.enabled[0] = false;
            //    tile.enabled[1] = false;
            //    tile.enabled[2] = false;
            //} else if column == 6 { //on right edge
            //    tile.enabled[3] = false;
            //    tile.enabled[4] = false;
            //    tile.enabled[5] = false;
            //} else if row == 0 { //first in column
            //    tile.enabled[2] = false;
            //    tile.enabled[5] = false;
            //} else if row == 5 { //topping out column
            //    tile.enabled[0] = false;
            //    tile.enabled[3] = false;
            //}

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

            // update upleft
            if search[0] {
                if let Some(tile_ul) = self.tiles[column-1][row+1] {
                    if tile_ul.color != color { //opposite color tile
                        
                        // disable search from opposite side
                        self.tiles[column - (1 + tile_ul.chains[0] as usize)][row + (1 + tile_ul.chains[0] as usize)].unwrap().update[5] = false; //safe unwrap

                    } else if search[5] { //same color tile, want to double check the other side for the same color, form a mega chain
                        match self.tiles[column+1][row-1] {
                            None => {
                                search[5] = false;
                            },
                            Some(tile_dr) => {
                                if tile_dr.color == color {
                                    
                                    // unify both ends of the chain by setting update and dist for both far ends
                                    let dist = 2 + tile_ul.chains[0] + tile_dr.chains[5]; //safe unwrap

                                    let tile_ul_far = &mut self.tiles[column - (1 + tile_ul.chains[0] as usize)][row + (1 + tile_ul.chains[5] as usize)].unwrap(); //safe unwrap
                                    let tile_dr_far = &mut self.tiles[column + (1 + tile_dr.chains[0] as usize)][row - (1 + tile_dr.chains[5] as usize)].unwrap(); //safe unwrap

                                    tile_ul_far.chains[5] = dist;
                                    tile_ul_far.update[5] = tile_dr.update[5];

                                    tile_dr_far.chains[0] = dist;
                                    tile_dr_far.update[0] = tile_ul.update[0];

                                    // disable search for dr, since we already handled the behavior
                                    search[5] = false;
                                }
                                // else tile is of opposite color, no special behavior here
                            },

                        }
                    } else { //same tile color, but other side is disabled
                        todo!();
                    }
                }
                // disabled top left, do nothing here
            }


            todo!()
        }

        // this column is completely filled up
        Event::Invalid
    }
}





