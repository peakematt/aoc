use std::collections::HashMap;

#[derive(Copy, Clone)]
struct Location {
    x: i32,
    y: i32,
    visit_count: i32,
}

impl Location {
    fn new(x: i32, y: i32) -> Location {
        Location {
            x,
            y,
            visit_count: 0,
        }
    }

    fn visit(&mut self) {
        self.visit_count += 1
    }

    fn hash_key(&self) -> String {
        format!("{},{}", self.x, self.y).to_string()
    }
}

struct Visitor {
    current_location: Location,
    visited_locations: HashMap<String, Location>,
}

impl Visitor {
    fn new() -> Visitor {
        let first_location = Location::new(0, 0);
        let mut b = Visitor {
            current_location: first_location,
            visited_locations: HashMap::new(),
        };
        b.visited_locations
            .insert(first_location.hash_key(), first_location);
        b
    }

    fn visit(&mut self, loc: Location) {
        match self.visited_locations.get_mut(&loc.hash_key()) {
            Some(location) => location.visit(),
            None => drop(
                self.visited_locations
                    .insert(loc.hash_key(), Location::new(loc.x, loc.y)),
            ),
        }
        self.current_location = loc;
    }

    fn handle_move(&mut self, direction: std::string::String) {
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
    }
}

fn main() {
    let mut santa = Visitor::new();
    let mut robo_santa = Visitor::new();
    let mut santa_turn = true;
    let mut all_locations: HashMap<String, Location> = HashMap::new();
    let input = std::fs::read_to_string("src/input.txt").expect("Should be able to read input");

    for c in input.chars() {
        if santa_turn {
            santa.handle_move(c.to_string());
            santa_turn = false;
        } else {
            robo_santa.handle_move(c.to_string());
            santa_turn = true;
        }
    }

    for k in santa.visited_locations.keys() {
        if all_locations.contains_key(k) {
            // pass
        } else {
            let v = santa
                .visited_locations
                .get(k)
                .expect("Hashmap contains key but no value at that key");
            all_locations.insert(k.to_owned(), Location::new(v.x, v.y));
        }
    }
    for k in robo_santa.visited_locations.keys() {
        if all_locations.contains_key(k) {
            // pass
        } else {
            let v = match robo_santa.visited_locations.get(k) {
                Some(location) => location,
                None => panic!("failed to get santa's visited location"),
            };
            all_locations.insert(k.to_owned(), Location::new(v.x, v.y));
        }
    }

    println!("Visited {} locations", all_locations.len())
}
