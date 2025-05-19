<script setup>

const emit = defineEmits(['escolha'])

function escolher(view) {
    emit('escolha', view)
}

</script>

<template>
    <div class="table-container">
        <table class="escolha-algum">
            <thead>
                <tr>
                    <th>Escolha:</th>
                </tr>
            </thead>
            <tbody>
                <tr @click="escolher('note')">
                    <td class="no-link">Note</td>
                </tr>
                <tr @click="escolher('checklist')">
                    <td class="no-link">Check-List</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>


<style scoped>

.table-container {
    margin: 20px;
}

.escolha-algum {
    border: 2px solid #ddd;
    font-family: Arial, sans-serif;
    text-align: center;
    border-radius: 10px; 
    border-spacing: 0; 
    overflow: hidden; 
}

.escolha-algum, .escolha-algum th, .escolha-algum td {
    height: 50px;
    width: 250px;
}

.escolha-algum th, .escolha-algum td {
    padding: 12px;
    border-bottom: 1px solid #ddd;
}

.escolha-algum th, .escolha-algum td {
    border-radius: 0;
}

.escolha-algum tbody tr:last-child td:last-child {
    border-bottom: 0;
}

.no-link {
  text-decoration: none; 
  color: inherit; 
  cursor: pointer; 
}

.escolha-algum td:hover {
    background: rgba(255, 255, 255, 0.204);
    cursor: pointer;
}

.escolha-algum thead th {
    background: rgba(0, 174, 255, 0.295);
}

</style>