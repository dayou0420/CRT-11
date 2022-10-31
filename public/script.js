const DEV_URL = 'https://crt-11.onrender.com/v1/tasks'
const LOCAL_URL = 'http://localhost:8080/v1/tasks'

async function myFetch(url) {
    const response = await fetch(url);
    const json = await response.json();
    return json;
}

(async function() {
    const f = await myFetch(DEV_URL);
    const date = f.data.data.map(m => m.date);
    const bill = f.data.data.map(m => m.bill);
    const powerBill = f.data.data.map(m => m.power.bill);
    const powerUsed = f.data.data.map(m => m.power.used);
    const gasBill = f.data.data.map(m => m.gas.bill);
    const gasUsed = f.data.data.map(m => m.gas.used);
    const lineData = {
        labels: date,
        datasets: [
            {
                label: 'Power and gas bill',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderColor: 'rgba(75, 192, 192, 1)',
                data: bill,
            },
            {
                label: 'Power bill',
                backgroundColor: 'rgba(255, 206, 86, 0.2)',
                borderColor: 'rgba(255, 206, 86, 1)',
                data: powerBill,
            },
            {
                label: 'Gas bill',
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                data: gasBill,
            }
        ]
    };
    const lineConfig = {
        type: 'line',
        data: lineData,
        options: {
            layout: {
                padding: 40
            }
        }
    };
    const myLineChart = new Chart(
        document.getElementById('myLineChart'),
        lineConfig
    );

    const linePowerUsedData = {
        labels: date,
        datasets: [
            {
                label: 'Power used',
                backgroundColor: 'rgba(255, 205, 86, 0.2)',
                borderColor: 'rgb(255, 205, 86)',
                data: powerUsed,
            }
        ]
    };
    const linePowerUsedConfig = {
        type: 'line',
        data: linePowerUsedData,
        options: {
            layout: {
                padding: 40
            }
        }
    };
    const myLinePowerUsedChart = new Chart(
        document.getElementById('myLinePowerUsedChart'),
        linePowerUsedConfig
    );

    const lineGasUsedData = {
        labels: date,
        datasets: [
            {
                label: 'Gas used',
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgb(255, 99, 132)',
                data: gasUsed,
            }
        ]
    };
    const lineGasUsedConfig = {
        type: 'line',
        data: lineGasUsedData,
        options: {
            layout: {
                padding: 40
            }
        }
    };
    const myLineGasUsedChart = new Chart(
        document.getElementById('myLineGasUsedChart'),
        lineGasUsedConfig
    );

    const billSum = bill.reduce((pre, cur) => pre + cur, 0);
    const powerBillSum = powerBill.reduce((pre, cur) => pre + cur, 0);
    const gasBillSum = gasBill.reduce((pre, cur) => pre + cur, 0);
    const labels = ["Power and gas bill", "Power bill", "Gas bill"]
    const data = {
    labels: labels,
    datasets: [{
        axis: 'y',
        label: 'Power and gas bill',
        data: [billSum,powerBillSum, gasBillSum],
        backgroundColor: [
        'rgba(75, 192, 192, 0.2)',
        'rgba(255, 205, 86, 0.2)',
        'rgba(255, 99, 132, 0.2)'
        ],
        borderColor: [
        'rgb(75, 192, 192)',
        'rgb(255, 205, 86)',
        'rgb(255, 99, 132)'
        ],
        borderWidth: 1
        }]
    };
    const config = {
        type: 'bar',
        data: data,
        options: {
            indexAxis: 'y',
            layout: {
                padding: 40
            }
        },
    };
    const myChart = new Chart(
        document.getElementById('myChart'),
        config
    );
})();
