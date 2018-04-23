extern crate iron;
extern crate router;

use iron::prelude::*;
use iron::status;
use router::Router;


fn main() {
    let mut router = Router::new();           // Alternative syntax:
    router.get("/", handler, "index");        // let router = router!(index: get "/" => handler,
    router.get("/:query", handler, "query");  //                      query: get "/:query" => handler);

    fn handler(req: &mut Request) -> IronResult<Response> {
        let ref query = req.extensions.get::<Router>().unwrap().find("query").unwrap_or("/");
        Ok(Response::with((status::Ok, *query)))
    }

    let _server = Iron::new(router).http("0.0.0.0:80").unwrap();
    println!("On 80");
}