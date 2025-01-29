const eventSource = new EventSource('/sse');
        //// No event filter
        // eventSource.onmessage = (event) => {
        //     const message = document.getElementById('messsage');
        //     message.innerHTML = event.data;
        //     console.log(event.data);
        // };

        // With event filter
        eventSource.addEventListener('messages', (event) => {
            const message = document.getElementById('messages');
            message.innerHTML = event.data;
            console.log(event.data);
        });