use std::collections::HashSet;

#[derive(Copy, Clone, Eq, PartialEq, Hash)]
struct Location {
    x: i32,
    y: i32,
}

struct Visitor {
    current_location: Location,
    visited_locations: HashSet<Location>,
}

impl Location {
    fn new(x: i32, y: i32) -> Location {
        Location { x, y }
    }
}

impl Visitor {
    fn new() -> Visitor {
        let first_location = Location::new(0, 0);
        let mut b = Visitor {
            current_location: first_location,
            visited_locations: HashSet::new(),
        };
        b.visit(first_location);
        b
    }

    fn visit(&mut self, loc: Location) {
        if !self.visited_locations.contains(&loc) {
            self.visited_locations.insert(loc);
        }
        self.current_location = loc;
    }

    fn handle_move(&mut self, direction: std::string::String) -> Location {
        match direction.as_ref() {
            "^" => self.visit(Location::new(
                self.current_location.x,
                self.current_location.y + 1,
            )),
            ">" => self.visit(Location::new(
                self.current_location.x + 1,
                self.current_location.y,
            )),
            "<" => self.visit(Location::new(
                self.current_location.x - 1,
                self.current_location.y,
            )),
            "v" => self.visit(Location::new(
                self.current_location.x,
                self.current_location.y - 1,
            )),
            "\n" => {}
            other => println!("Unknown input character: '{}'", other),
        }
        self.current_location
    }
}

fn main() {
    let mut santa = Visitor::new();
    let mut robo_santa = Visitor::new();
    let mut santa_turn = true;
    let mut all_locations: HashSet<Location> = HashSet::new();
    let input = std::fs::read_to_string("src/input.txt").expect("Should be able to read input");

    for c in input.chars() {
        if santa_turn {
            all_locations.insert(santa.handle_move(c.to_string()));
            santa_turn = false;
        } else {
            all_locations.insert(robo_santa.handle_move(c.to_string()));
            santa_turn = true;
        }
    }

    println!("Visited {} locations", all_locations.len());
}
