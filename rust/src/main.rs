extern crate iron;
extern crate router;

use iron::prelude::*;
use router::Router;

mod heat_controller;

fn main() {
    let mut router = Router::new();

    //set routes
    router.get("/userHeats/:query", heat_controller::get_all_user_has_heat_zone, "index");
    router.post("/heats", heat_controller::store, "query");

    let _server = Iron::new(router).http("0.0.0.0:80").unwrap();
    println!("On 80");
}