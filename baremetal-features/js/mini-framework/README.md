## Phase 1: The Router
    **Description: The router that swaps a static "Home" string for an "About" string in the #app div.**
    **Consept: In an SPA, we want to tell the browser: "Stay right here, don't refresh, just change the text in the URL bar and swap the HTML inside this specific div."**
    **Goal: Swap content based on the URL without a page refresh.**
    **The Flow:**
        Click link → Prevent Reload → Update URL bar → Find Match → Inject HTML.
        Back Button → Find Match → Inject HTML

    Task 1.1: Create an index.html with a <div id="app"></div> and a navigation menu using data-link attributes instead of standard href behavior.
    Task 1.2: Write a navigateTo(url) function that uses history.pushState to change the URL.
    Task 1.3: Create a router() function that looks at window.location.pathname, finds the matching component, and injects its HTML into the #app container.
    Task 1.4: Listen for the popstate event (back/forward buttons) and intercept click events on links to prevent page reloads.

## Phase 2: Proxy-based State (Reactivity)
    Description: Implement the Proxy-based State so changing a variable updates the text on the screen without manually calling document.getElementById.
    Goal: Make the UI "react" to data changes automatically.
    Task 2.1: Create a global state object.
    Task 2.2: Wrap this object in a new Proxy(state, {...}).
    Task 2.3: In the Proxy's set handler, add a trigger that calls the router() function every time a property is modified.
    Task 2.4: Update a component to display a state value (e.g., a counter) and add a button with an onclick that increments state.count.
    Note: Since you're re-rendering the whole string, you'll learn why "Focus" is lost on inputs—a problem React solves with the Virtual DOM.

## Phase 3: The "MiniJS" Abstraction
    Description: Abstract these into a class or library called something like MiniJS so you can initialize it with new MiniJS({ routes, state }).
    Goal: Package the logic into a reusable class.
    Task 3.1: Create a MiniJS class in app.js.
    Task 3.2: Define a constructor that accepts a configuration object: { target, routes, initialState }.
    Task 3.3: Move the Proxy logic and the Router logic inside the class methods.
    Task 3.4: Implement an init() method that attaches the event listeners and performs the initial render.
    Task 3.5: Refactor your index.html to simply call

## Declarative Events (The Parser)
## The Template "Compiler"
## Component Lifecycle

## structure 
mini-framework/
├── index.html          # The entry point
├── src/
│   ├── app.js          # Main framework logic (MiniJS class)
│   ├── routes.js       # Route definitions
│   └── components/     # Your view functions
│       ├── Home.js
│       └── About.js
└── style.css