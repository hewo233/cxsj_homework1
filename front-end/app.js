function queryData() {
    const input = document.getElementById('queryInput').value;
    fetch(`http://localhost:8080/query?email=${encodeURIComponent(input)}`)
        .then(response => response.json())
        .then(data => {
            const table = document.getElementById('dataTable').getElementsByTagName('tbody')[0];
            table.innerHTML = '';  // 清空现有的表格内容
            const tr = document.createElement('tr');
            // 将每个属性添加到表格行
            tr.appendChild(createCell(data.id));
            tr.appendChild(createCell(data.email));
            tr.appendChild(createCell(data.name));
            tr.appendChild(createCell(data.password)); // 注意: 密码通常不应显示
            tr.appendChild(createCell(data.gender));
            table.appendChild(tr);
        })
        .catch(error => console.error('Error:', error));
}

function createCell(text) {
    const td = document.createElement('td');
    td.textContent = text;
    return td;
}

// 添加数据
function addData() {
    const email = document.getElementById('addEmail').value;
    const name = document.getElementById('addName').value;
    const password = document.getElementById('addPassword').value;
    const gender = document.getElementById('addGender').value;

    const userData = {
        email: email,
        name: name,
        password: password,
        gender: gender
    };

    console.log(JSON.stringify(userData)); // 输出查看发送的数据是否正确

    fetch('http://localhost:8080/add', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(userData)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log('Data added successfully', data);
            alert('Data added successfully');
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Failed to add data');
        });
}




// 更新数据
function updateData() {
    const email = document.getElementById('updateEmail').value;
    const name = document.getElementById('updateName').value;
    const password = document.getElementById('updatePassword').value;
    const gender = document.getElementById('updateGender').value;

    const updateData = {
        email: email,
        name: name,
        password: password,
        gender: gender
    };

    console.log(JSON.stringify(updateData)); // 输出查看发送的数据是否正确

    fetch('http://localhost:8080/modify', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(updateData)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log('Data added successfully', data);
            alert('Data added successfully');
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Failed to add data');
        });
}

function deleteData() {
    const email = document.getElementById('deleteEmail').value;

    fetch(`http://localhost:8080/delete?email=${encodeURIComponent(email)}`, {
        method: 'DELETE'
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            console.log('Data deleted successfully');
            alert('Data deleted successfully');
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Failed to delete data');
        });
}


