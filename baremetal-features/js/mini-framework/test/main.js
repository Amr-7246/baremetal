import MiniJS from "../src/app"
import {Home} from "./Home"

const routes = [
    {path: '/home', view: Home},
    { path: "/", view: () => "<h1>Welcome Home</h1><p>This is the landing page.</p>" },
    { path: "/about", view: () => "<h1>About Us</h1><p>We are building a framework!</p>" },
    { path: "/contact", view: () => "<h1>Contact</h1><p>Email us at dev@mini.js</p>" },
    { path: "/404", view: () => "<h1>404</h1><p>Page not found</p>" }
]
const initialState = {
    count:0
}
const app = new MiniJS({target: '#app', routes: routes, initialState: initialState})
app.init()