import {router, navigateTo} from "./router.js"

// in order to make a single page application we need routing
// the routing is not just interacting with just buttons and direct referehes
// we need to make sure
//     we allow the movement of links normally
//         if you click on home you move /home
//         browser buttons back and forward should work

// and for that we need a router. 
// the router detects any change in the current link 
// and also detects any click on a button og interest


//listens for any change in the url (popstate). used for forward and backward buttons in browser
window.addEventListener("popstate", router);


// checking for any click in the window. this is better than making a click for every single event.
// why? when we change the pages, some buttons will disappear so we will needd to make another event listener once again.
document.addEventListener("click", (e)=> {

    //checking for an a href with data-link field
    const link = e.target.closest("a[data-link]")
    if(link){
        e.preventDefault();
        navigateTo(link.href);
        return;
    }

    //checking for a button with daata-action field
    const btn = e.target.closest("[data-action]");
    if(btn) {
        const action = btn.dataset.action;

        //basically we can use this if we still need a button
        // in the beginnign we wont but as we move up and intiate the chats we will definitely will
        console.log("action:", action);
        if (action =="Chat") {
            navigateTo("/chat")
        } else if( action =="Create") {
            navigateTo("/create")
        }
    }
});


//checking for form submissions here
document.addEventListener("submit", (e) => {
    const form = e.target.closest('[data-form]')
    if(!form) return;
    
    e.preventDefault();

    const formName = form.dataset.from;
    const data = Object.fromEntries(new FormData(form))

     console.log("form:", formName, data);
})

router();