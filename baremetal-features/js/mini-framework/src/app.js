// @ts-nocheck
//TODO: check the another frameworks manner to improv that simple version
export default class MiniJS {
    constructor({target, routes, initialState}){
        this.target = document.querySelector(target)
        this.routes = routes
        //& Make the app "reacting", global state + Proxy
            //TODO: solve that problem ... state change → destroy everything → rebuild everything 
            this.state = new Proxy(initialState, { //! this state which will be used in changing/access the data .. state.count++ || state.count
                set: (target, prop, value) => {
                    target[prop] = value
                    this.render() // re-mount the UI
                    return true
                }
            })
    }
    //& match the routes and inject its content
        render() {
            const match = this.routes.find(r => r.path === location.pathname) || this.routes.find(r => r.path === "/404")
            this.target.innerHTML = match.view(this.state) //! inject the global state  at the view
        }
    //& intercept the clicks
        init(){
            window.addEventListener("popstate", () => this.render()) //! here the arrow function over Regular Function to insure that this refare to the MiniJS class not the browser window
            //! we do not need the "DOMContentLoaded" event here because of we importing the JS script in a module mode which do the same effect 
            document.body.addEventListener("click", (e) => {
                if(e.target.matches("[data-link]")){ //! using the CSS selector function over hasAttribute() is more powerfull
                    e.preventDefault()
                    history.pushState(null, "", e.target.href) //! change the route at the browser this.render bar without refreshing the page 
                }
            })
            this.render() //! here we inject the initial page content
        }
}