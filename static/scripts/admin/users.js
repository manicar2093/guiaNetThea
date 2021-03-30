document.addEventListener("DOMContentLoaded", () => {

    const listenEventModalFromModalBackground = (modalBackground, modal) => {
        modalBackground.addEventListener("click", () => {
            modal.classList.toggle("is-active")
        })
    }

    const listenEventsButtons = (buttons, modal) => {
        buttons.forEach(i => {
            i.addEventListener("click", () => {
                modal.classList.toggle("is-active")
            })
        })
    }
    
    let restore_buttons = Array.from(document.querySelectorAll(".is-warning"))
    let modal = document.getElementById("restore_password_form")
    listenEventsButtons(restore_buttons, modal)

    let modalBackground = document.querySelector("#restore_password_form .modal-background")
    listenEventModalFromModalBackground(modalBackground, modal)

    let delete_buttons = Array.from(document.querySelectorAll(".is-danger"))
    let modalDelete = document.getElementById("delete_user_form")
    listenEventsButtons(delete_buttons, modalDelete)

    let modalBackgroundDelete = document.querySelector("#delete_user_form .modal-background")
    listenEventModalFromModalBackground(modalBackgroundDelete, modalDelete)


})