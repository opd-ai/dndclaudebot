// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"github.com/a-h/templ"
	templruntime "github.com/a-h/templ/runtime"
)

import (
	"fmt"
	"log"
)

// Helper function to log and return empty string
func logf(format string, args ...interface{}) string {
	log.Printf(format, args...)
	return ""
}

// Helper function to log simple messages and return empty string
func logln(msg string) string {
	log.Println(msg)
	return ""
}

func GeneratorForm() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"generator-container\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(logln("Rendering GeneratorForm"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `srv/components/generator.templ`, Line: 22, Col: 42}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"generator-form\" hx-post=\"/generate\" hx-target=\"#generation-status\"><textarea name=\"prompt\" required placeholder=\"Enter your prompt...\"></textarea> <button type=\"submit\">Generate</button></form><div id=\"generation-status\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = GenerationStatus("").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func wsOutputContainer(sessionID string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(logf("Rendering wsOutputContainer for session: %s", sessionID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `srv/components/generator.templ`, Line: 34, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"ws-output\" class=\"output-container\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/api/messages/%s", sessionID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `srv/components/generator.templ`, Line: 38, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"load delay:100ms, every 2s\" hx-swap=\"innerHTML\"><div class=\"loading\">Loading messages...</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func wsStyles() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(logln("Applying WebSocket styles"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `srv/components/generator.templ`, Line: 47, Col: 40}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n        .output-container {\n            min-height: 100px;\n            max-height: 400px;\n            overflow-y: auto;\n            padding: 1rem;\n            border: 1px solid #ccc;\n            margin-top: 1rem;\n            background: #f9f9f9;\n            scroll-behavior: smooth;\n        }\n        .message {\n            margin-bottom: 0.5rem;\n            padding: 0.5rem;\n            border-left: 3px solid #ccc;\n            background: white;\n            animation: fadeIn 0.3s ease-out;\n        }\n        @keyframes fadeIn {\n            from { opacity: 0; transform: translateY(5px); }\n            to { opacity: 1; transform: translateY(0); }\n        }\n        .message.generating { border-color: #ffd700; }\n        .message.completed { border-color: #4caf50; }\n        .message.error { border-color: #f44336; }\n        .loading {\n            text-align: center;\n            color: #666;\n            padding: 1rem;\n        }\n        .message-header {\n            display: flex;\n            justify-content: space-between;\n            align-items: center;\n            margin-bottom: 0.5rem;\n            font-size: 0.9em;\n        }\n        .timestamp {\n            font-size: 0.8em;\n            color: #666;\n        }\n        .message pre {\n            white-space: pre-wrap;\n            word-break: break-word;\n            background: #f5f5f5;\n            padding: 0.5rem;\n            border-radius: 4px;\n            margin: 0.5rem 0;\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func wsHandling() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_wsHandling_bc89`,
		Function: `function __templ_wsHandling_bc89(){console.log('Initializing WebSocket handling script');
    
    const state = {
        messageCache: new Set(),
        isConnected: false,
        sessionId: document.querySelector('[data-session-id]')?.dataset.sessionId
    };
    console.log('Initial state:', { 
        cacheSize: state.messageCache.size, 
        isConnected: state.isConnected, 
        sessionId: state.sessionId 
    });

    function initializeMessageHandling() {
        console.log('Initializing message handling');
        const output = document.getElementById('ws-output');
        if (!output || !state.sessionId) {
            console.warn('Missing output element or session ID');
            return;
        }
        loadMessages();
    }

    function loadMessages() {
        console.log('Loading messages for session:', state.sessionId);
        const sessionId = state.sessionId;
        if (!sessionId) {
            console.warn('No session ID available');
            return;
        }

        fetch(` + "`" + `/api/messages/${sessionId}` + "`" + `)
            .then(response => {
                console.log('Message fetch response status:', response.status);
                return response.json();
            })
            .then(messages => {
                console.log('Received messages:', messages.length);
                const output = document.getElementById('ws-output');
                if (!output) {
                    console.warn('Output element not found');
                    return;
                }
                output.innerHTML = '';
                state.messageCache.clear();
                messages.forEach(addMessageToOutput);
            })
            .catch(error => {
                console.error('Error loading messages:', error);
            });
    }

    function formatTimestamp(timestamp) {
        console.debug('Formatting timestamp:', timestamp);
        return new Date(timestamp).toLocaleTimeString(undefined, {
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit'
        });
    }

    function getMessageId(data) {
        const id = ` + "`" + `${data.timestamp}-${data.status}-${data.message}` + "`" + `;
        console.debug('Generated message ID:', id);
        return id;
    }

    function addMessageToOutput(data) {
        console.log('Adding message to output:', data);
        const output = document.getElementById('ws-output');
        if (!output) {
            console.warn('Output element not found');
            return;
        }

        const messageId = getMessageId(data);
        if (state.messageCache.has(messageId)) {
            console.debug('Message already exists:', messageId);
            return;
        }
        state.messageCache.add(messageId);
        console.debug('Cache size after add:', state.messageCache.size);

        const div = document.createElement('div');
        div.className = ` + "`" + `message ${data.status || 'generating'}` + "`" + `;
        div.setAttribute('data-message-id', messageId);
        
        div.innerHTML = ` + "`" + `
            <div class="message-header">
                <strong>${data.status || 'generating'}</strong>
                <span class="timestamp">${formatTimestamp(data.timestamp)}</span>
            </div>
            ${data.message ? ` + "`" + `<p>${data.message}</p>` + "`" + ` : ''}
            ${data.output ? ` + "`" + `<pre>${data.output}</pre>` + "`" + ` : ''}
        ` + "`" + `;
        
        output.appendChild(div);
        console.log('Message added to DOM:', messageId);
        
        requestAnimationFrame(() => {
            output.scrollTop = output.scrollHeight;
            console.debug('Scrolled to bottom');
        });
    }

    document.addEventListener('DOMContentLoaded', () => {
        console.log('DOM Content Loaded - initializing message handling');
        initializeMessageHandling();
    });

    document.addEventListener('htmx:wsAfterMessage', function(evt) {
        console.log('WebSocket message received:', evt.detail);
        try {
            const data = JSON.parse(evt.detail.message);
            console.log('Parsed WebSocket message:', data);
            addMessageToOutput(data);
        } catch (error) {
            console.error('Error processing WebSocket message:', error);
            loadMessages();
        }
    });

    ['wsConnecting', 'wsOpen', 'wsClose', 'wsError'].forEach(event => {
        document.addEventListener(` + "`" + `htmx:${event}` + "`" + `, (evt) => {
            console.log(` + "`" + `WebSocket ${event}:` + "`" + `, evt.detail);
            if (event === 'wsOpen') {
                state.isConnected = true;
                console.log('WebSocket connected');
                loadMessages();
            } else if (event === 'wsClose' || event === 'wsError') {
                state.isConnected = false;
                console.log('WebSocket disconnected or error');
                setTimeout(loadMessages, 1000);
            }
        });
    });

    document.addEventListener('visibilitychange', function() {
        console.log('Visibility changed:', !document.hidden);
        if (!document.hidden) {
            loadMessages();
        }
    });

    htmx.on('#ws-output', 'htmx:afterRequest', function(evt) {
        console.log('HTMX after request:', evt.detail.xhr.status);
        if (evt.detail.xhr.status === 200) {
            try {
                const messages = JSON.parse(evt.detail.xhr.responseText);
                console.log('Received messages in HTMX request:', messages.length);
                const output = document.getElementById('ws-output');
                if (!output) {
                    console.warn('Output element not found');
                    return;
                }
                output.innerHTML = '';
                state.messageCache.clear();
                messages.forEach(addMessageToOutput);
            } catch (error) {
                console.error('Error parsing messages:', error);
            }
        }
    });
}`,
		Call:       templ.SafeScript(`__templ_wsHandling_bc89`),
		CallInline: templ.SafeScriptInline(`__templ_wsHandling_bc89`),
	}
}

func GenerationStatus(sessionID string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(logf("Rendering GenerationStatus for session: %s", sessionID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `srv/components/generator.templ`, Line: 267, Col: 67}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if sessionID != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"ws-status\" hx-ext=\"ws,json-enc\" ws-connect=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/ws/%s", sessionID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `srv/components/generator.templ`, Line: 272, Col: 57}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"generation-status\" data-session-id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(sessionID)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `srv/components/generator.templ`, Line: 274, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = wsOutputContainer(sessionID).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = wsStyles().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = wsHandling().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(logln("Warning: Empty sessionID in GenerationStatus"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `srv/components/generator.templ`, Line: 281, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate