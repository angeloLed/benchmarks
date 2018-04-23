extern crate iron;
extern crate router;
extern crate bodyparser;

use iron::prelude::*;
use iron::status;
use router::Router;

pub fn get_all_user_has_heat_zone(req: &mut Request) -> IronResult<Response> {
    let ref query = req.extensions.get::<Router>().unwrap().find("query").unwrap_or("/");
    Ok(Response::with((status::Ok, *query)))
}

pub fn store(req: &mut Request) -> IronResult<Response> {
    let ref query = req.extensions.get::<Router>().unwrap().find("query").unwrap_or("/");
    Ok(Response::with((status::Ok, *query)))
}