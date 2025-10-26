use std::collections::HashMap;

#[derive(Debug)]
pub struct City {
    pub name: String,
    pub routes: HashMap<String, u32>, // neighbor name -> cost
}

impl City {
    pub fn new(name: &str) -> Self {
        City {
            name: name.to_string(),
            routes: HashMap::new(),
        }
    }

    pub fn connect(&mut self, to: &mut City, cost: u32) {
        self.routes.insert(to.name.clone(), cost);
        to.routes.insert(self.name.clone(), cost); // bidirectional
    }
}

pub fn dijkstra_shortest_path(
    cities: &HashMap<String, City>,
    start: &str,
    goal: &str,
) -> Option<(u32, Vec<String>)> {
    let mut dist: HashMap<String, u32> = HashMap::new();
    let mut prev: HashMap<String, String> = HashMap::new();
    let mut unvisited: Vec<String> = cities.keys().cloned().collect();

    // set start distance to 0
    for city in cities.keys() {
        dist.insert(city.clone(), if city == start { 0 } else { u32::MAX });
    }

    while !unvisited.is_empty() {
        // find unvisited city with smallest known distance
        let current = unvisited
            .iter()
            .min_by(|a, b| {
                let da = dist.get(*a).unwrap_or(&u32::MAX);
                let db = dist.get(*b).unwrap_or(&u32::MAX);
                da.cmp(db)
            })
            .cloned()
            .unwrap();

        if current == goal {
            break;
        }

        unvisited.retain(|x| x != &current);

        if let Some(city) = cities.get(&current) {
            let current_dist = *dist.get(&current).unwrap_or(&u32::MAX);

            for (neighbor, &cost) in &city.routes {
                if !unvisited.contains(neighbor) {
                    continue;
                }
                let alt = current_dist.saturating_add(cost);
                if alt < *dist.get(neighbor).unwrap_or(&u32::MAX) {
                    dist.insert(neighbor.clone(), alt);
                    prev.insert(neighbor.clone(), current.clone());
                }
            }
        }
    }

    // reconstruct path
    let mut path = Vec::new();
    let mut current = goal.to_string();
    while let Some(p) = prev.get(&current) {
        path.push(current.clone());
        current = p.clone();
    }
    path.push(start.to_string());
    path.reverse();

    dist.get(goal).copied().map(|d| (d, path))
}
