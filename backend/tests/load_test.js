import { check } from 'k6';
import http from 'k6/http';
import { htmlReport } from 'https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js';

export let options = {
    stages: [
        { duration: '30s', target: 50 },
        { duration: '30s', target: 100 },
        { duration: '30s', target: 150 },
        { duration: '30s', target: 200 },
        { duration: '30s', target: 250 },
        { duration: '30s', target: 300 },
        { duration: '30s', target: 350 },
    ],
    thresholds: {
        http_req_duration: ['p(95)<500'],
    },
};

// Usando a variável de ambiente K6_PORT para definir a porta
const port = __ENV.K6_PORT || 8080; // Porta padrão caso K6_PORT não esteja definida

export default function () {

    let payload = JSON.stringify({
        "g-recaptcha-response": "",
        "comment": "Apenas um comentário",
        "name": "Marcelo",
        "mail": "my.name@email.com"
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    let res = http.post(`http://app:${port}/contact`, payload, params);

    check(res, {
        'status is 201': (r) => r.status === 201,
        'duration < 500ms': (r) => r.timings.duration < 500,
    });
}

export function handleSummary(data) {
    // Logar as métricas no console para referência
    console.log('Resumo dos testes:');
    console.log(JSON.stringify(data, null, 2));
    return {
        '/tests/result.html': htmlReport(data), // Força o caminho do relatório HTML para ser salvo no diretório /tests
    };
}
