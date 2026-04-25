//& Global state 
    let currentUser = ''
    let lastMessageId = 0;
    let isPolling = false;

//& DOM Elements
    const loginContainer = document.getElementById('login-container');
    const chatContainer = document.getElementById('chat-container');
    const loginForm = document.getElementById('login-form');
    const messageForm = document.getElementById('message-form');
    const messagesContainer = document.getElementById('messages-container');
    const messageInput = document.getElementById('message-input');
    const currentUsernameSpan = document.getElementById('current-username');
    const loginError = document.getElementById('login-error');

//& login logic
    loginForm.addEventListener("submit", async (e) => {
        e.preventDefault()
        const username = document.getElementById('username').value.trim();
        
        if (!username) {
            showLoginError('Please enter a username');
            return;
        }
        
        //~ Disable login button during request
        const submitBtn = loginForm.querySelector('button');
        submitBtn.disabled = true;
        submitBtn.textContent = 'Joining...';

        //~ Send data to server
        try {
            const response = await fetch('api/login.php', {
                method: 'POST',
                headers: 'Content-Type : application/json',
                body: username
            })
            const data = await response.json()
            if (response.ok && data.success) {
                currentUser = data.user
                showChatUI();
                loadHistory();
            } else {
                showLoginError(data.error || 'Login failed');
                submitBtn.disabled = false;
                submitBtn.textContent = 'Join Chat';
            }
        } catch (error) {
            console.error('Login error:', error);
            showLoginError('Network error. Please try again.');
            submitBtn.disabled = false;
            submitBtn.textContent = 'Join Chat';
        }
    })
//& pulling history logic
    async function loadHistory() {
        try {
            const response = await fetch('api/history.php');
            const messages = await response.json();
            
            if (response.ok) {
                messagesContainer.innerHTML = '';
                
                if (messages.length === 0) {
                    showEmptyState();
                } else {
                    messages.forEach(message => {
                        appendMessageToUI(message);
                    });
                    
                    lastMessageId = messages[messages.length - 1].id;
                    scrollToBottom();
                }
            } else {
                console.error('Failed to load history:', messages.error);
                showErrorMessage('Failed to load message history');
            }
        } catch (error) {
            console.error('History error:', error);
            showErrorMessage('Network error loading history');
        }
    }
//& Helper function 
    function showLoginError(message) {
        loginError.textContent = message;
        setTimeout(() => {
            loginError.textContent = '';
        }, 3000);
    }

    function showChatUI() {
        loginContainer.classList.add('hidden');
        chatContainer.classList.remove('hidden');
        currentUsernameSpan.textContent = currentUser.username;
        messageInput.focus();
    }

    function appendMessageToUI(message) {
        const messageDiv = document.createElement('div');
        messageDiv.className = `message ${message.username === currentUser.username ? 'own-message' : 'other-message'}`;
        messageDiv.setAttribute('data-message-id', message.id);
        
        const usernameSpan = document.createElement('div');
        usernameSpan.className = 'message-username';
        usernameSpan.textContent = escapeHtml(message.username);
        
        const textSpan = document.createElement('div');
        textSpan.className = 'message-text';
        textSpan.textContent = escapeHtml(message.message);
        
        const timeSpan = document.createElement('div');
        timeSpan.className = 'message-time';
        timeSpan.textContent = formatTime(message.created_at);
        
        messageDiv.appendChild(usernameSpan);
        messageDiv.appendChild(textSpan);
        messageDiv.appendChild(timeSpan);
        
        messagesContainer.appendChild(messageDiv);
    }
    function showEmptyState() {
        const emptyDiv = document.createElement('div');
        emptyDiv.className = 'empty-state';
        emptyDiv.innerHTML = `
            <p>No messages yet</p>
            <small>Be the first to send a message!</small>
        `;
        messagesContainer.appendChild(emptyDiv);
    }

    function showErrorMessage(message) {
        const errorDiv = document.createElement('div');
        errorDiv.className = 'error-state';
        errorDiv.textContent = message;
        messagesContainer.appendChild(errorDiv);
        
        setTimeout(() => {
            errorDiv.remove();
        }, 3000);
    }

    function scrollToBottom() {
        messagesContainer.scrollTop = messagesContainer.scrollHeight;
    }

    function formatTime(timestamp) {
        const date = new Date(timestamp);
        const now = new Date();
        const diff = now - date;
        
        // If less than 1 minute, show "Just now"
        if (diff < 60000) {
            return 'Just now';
        }
        
        // If less than 1 hour, show minutes ago
        if (diff < 3600000) {
            const minutes = Math.floor(diff / 60000);
            return `${minutes} minute${minutes > 1 ? 's' : ''} ago`;
        }
        
        // Otherwise show time
        return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
    }

    function escapeHtml(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }