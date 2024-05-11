////////
//Get data and populate the page
////////

document.addEventListener('DOMContentLoaded', onPageLoad);

async function onPageLoad(){
    const backendUrl = 'http://localhost:8080/api'; //todo: get from env

    const response = await fetch(backendUrl + "/contributions")
    const json_contributions = await response.json()

    beerDiv = document.querySelector('#collapseBeer')
    beerDiv.innerHTML = ""
    beerQuantity = 0

    json_contributions.forEach(json_contribution => {   //todo: make this for every drink dynamicly
        if(json_contribution.type === 'beer'){
            var div = document.createElement('div')
            div.classList.add('card', 'card-body', 'custom-collapse-3')
            div.innerHTML = json_contribution.contributor_name + ': ' + json_contribution.quantity + "l, "/*todo, make this dynamic*/ + json_contribution.description
            beerDiv.appendChild(div)
            beerQuantity += json_contribution.quantity 
        }
    });

    document.querySelector('beerText').innerText = 'Sör: ' + beerQuantity + 'l' //todo: if 0 ---> put ':(' behind
    console.log(beerDiv);
    console.log(json_contributions)
}
console.log("alma")

////////
//js stop event propagation
////////
// fix this later
// useful link: https://stackoverflow.com/questions/16914747/prevent-bootstrap-collapse-from-collapsing

// no_event_propagation_elements = document.querySelectorAll(".custom-js-no-event-propagation")
// console.log(no_event_propagation_elements);
// no_event_propagation_elements.forEach(element => {

//     element.addEventListener('click', function(event) {
//         event.preventDefault()
//         console.log("its here");
//         event.stopPropagation();
//         var currentAccordionButton = this.parentElement;
//         var ariaControls = currentAccordionButton.getAttribute("aria-controls");
//         // var headerAccordion = currentAccordionButton.parentElement;
//         // var containerAccordion = headerAccordion.parentElement;
//         // var accordionDiv = containerAccordion.querySelector("#" + ariaControls);
//         var accordionDiv = document.querySelector("#" + ariaControls);
//         accordionDiv.classList.add("show");
//         currentAccordionButton.setAttribute("aria-expanded", "true");
//         currentAccordionButton.classList.remove("collapsed");
//     });
// });
// debug = document.querySelector("#collapseDrinks")
// debug.addEventListener('click', function(){
//     console.log(debug)
//     console.log("parent clicked")
// })



////////
//modal js
////////
var exampleModal = document.getElementById('drinksModal')
exampleModal.addEventListener('show.bs.modal', function (event) {
  // Button that triggered the modal
  var button = event.relatedTarget
  // Extract info from data-bs-* attributes
  var type = button.getAttribute('data-bs-whatever')
  // If necessary, you could initiate an AJAX request here
  // and then do the updating in a callback.
  //
  // Update the modal's content.
  var modalBodyInput = exampleModal.querySelector('.modal-body select')

  modalBodyInput.value = type
})