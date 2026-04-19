// @ts-nocheck
//TODO: check the another frameworks manner to improv that simple version
//& change the route at the browser router bar without refreshing the page 
    const navigateTo = (url) => { 
        history.pushState(null, "", url)
        router()
    }
//& match the routes and inject its content
    const router = () => {
        //TODO: these routes should be dynamicly declared
        const routes = [
            { path: "/", view: () => "<h1>Welcome Home</h1><p>This is the landing page.</p>" },
            { path: "/about", view: () => "<h1>About Us</h1><p>We are building a framework!</p>" },
            { path: "/contact", view: () => "<h1>Contact</h1><p>Email us at dev@mini.js</p>" },
            { path: "/404", view: () => "<h1>404</h1><p>Page not found</p>" }
        ];
        const match = routes.find(r => r.path === location.pathname) || routes.find(r => r.path === "/404")
        document.getElementById("app")?.innerHTML = match.view()
    }
//& intercept the clicks
    window.addEventListener("popstate", router) 
    document.addEventListener('DOMContentLoaded', () => {
        document.body.addEventListener("click", (e) => {
            if(e.target.matches("[data-link]")){ //! using the CSS selector function over hasAttribute() is more powerfull
                e.preventDefault()
                navigateTo(e.target.href)
            }
        })
        router() //! here we inject the initial page content
    })
