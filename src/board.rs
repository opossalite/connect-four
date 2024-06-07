


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
    pub enabled: [bool; 6],
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
                enabled: [true, true, true, true, true, true],
            });
            let tile = &mut cur_col[row].unwrap();

            // cover some edge cases for enabled-ness
            if column == 0 { //on left edge
                tile.enabled[0] = false;
                tile.enabled[1] = false;
                tile.enabled[2] = false;
            } else if column == 6 { //on right edge
                tile.enabled[3] = false;
                tile.enabled[4] = false;
                tile.enabled[5] = false;
            } else if row == 0 { //first in column
                tile.enabled[2] = false;
                tile.enabled[5] = false;
            } else if row == 5 { //topping out column
                tile.enabled[0] = false;
                tile.enabled[3] = false;
            }

            // update chain lengths
            todo!()
        }

        // this column is completely filled up
        Event::Invalid
    }
}





