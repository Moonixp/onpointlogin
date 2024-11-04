fetch("http://localhost:8000/users")
    .then(response => response.json())
    .then(data => {
        const tbody = document.getElementById('nameRows');
        data.forEach(name => {
            const row = document.createElement('tr');
            row.innerHTML = `
        <td class="border border-[#FCFCFF] font-bold px-4 py-2">${name.fullname}</td>
        <td class="border border-[#FCFCFF] px-4 py-2">
            <button 
                onclick="Login('${name.fullname}')"
                class="bg-[#003366] border-solid border-2 border-[#001a33] px-5 py-1  text-[#FCFCFF] rounded-lg">
                Login
            </button>
        </td>
    `;
            tbody.appendChild(row);
        });
    })

function filterNames() {
    const input = document.getElementById('searchInput');
    const filter = input.value.toLowerCase();
    const rows = tbody.getElementsByTagName('tr');

    for (let i = 0; i < rows.length; i++) {
        const nameCell = rows[i].getElementsByTagName('td')[0].textContent.toLowerCase();
        rows[i].style.display = nameCell.includes(filter) ? '' : 'none';
    }
}

/**
 * Fetches the user with the given name from the server and logs them in.
 * The response is handled as follows:
 * - If the user is found, the user is logged in and a success message is displayed.
 * - If the user was already logged in, a message indicating that the user is already logged in is displayed.
 * - If the user is not found, a message indicating that the user was not found is displayed.
 * The messages are displayed for 1 second before the welcome message is reset.
 * @param {string} name - The name of the user to log in.
 */
const Login = (name) => {
    fetch("http://localhost:8000/loginuser", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name: name })
    })
        .then(response => {
            if (response.status === 200)
                return response.json()
            else
                return {}
        })
        .then(data => {

            if (data.id === undefined) {
                document.getElementById('welcome').innerText = "failed to load user, Try again";
            }
            else if (data.id !== undefined) {
                if (data.id === -1 && data.data !== undefined) {
                    document.getElementById('welcome').innerText = `${name} has already logged in`;
                } else {
                    document.getElementById('welcome').innerText = `Welcome ${name}`;
                }
            }

            setTimeout(function() {
                document.getElementById('welcome').innerText = "Welcome to OnPoint";
            }, 1000);

        })
}
