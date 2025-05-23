package microapidoc

import (
	"github.com/gin-gonic/gin"
)

func (c *Microapidoc) DocIndexHAndler(ctx *gin.Context) {

	indexHtml := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>API Documentation</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
        }
        header {
            background-color: blue;
            color: white;
            padding: 20px;
            margin-bottom: 20px;
            border-radius: 4px;
        }
        .auth-section {
            background-color: #f8f8f8;
            border: 1px solid #ddd;
            border-radius: 4px;
            padding: 15px;
            margin-bottom: 20px;
        }
        .auth-title {
            font-weight: bold;
            margin-bottom: 10px;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .auth-table {
            width: 100%;
            border-collapse: collapse;
        }
        .auth-table th, .auth-table td {
            text-align: left;
            padding: 8px;
            border-bottom: 1px solid #eee;
        }
        .auth-table th {
            background-color: #f5f5f5;
            font-weight: normal;
            color: #666;
        }
        .auth-table input {
            width: 100%;
            padding: 6px;
            border: 1px solid #ddd;
            border-radius: 3px;
            font-family: monospace;
        }
        .auth-saved-indicator {
            display: none;
            color: #4CAF50;
            font-size: 0.8em;
            margin-left: 10px;
            animation: fadeOut 2s forwards;
        }
        .auth-default {
            font-size: 0.8em;
            color: #666;
        }
        @keyframes fadeOut {
            0% { opacity: 1; }
            70% { opacity: 1; }
            100% { opacity: 0; }
        }
        .group {
            border: 1px solid #ddd;
            border-radius: 4px;
            margin-bottom: 20px;
            overflow: hidden;
        }
        .group-header {
            background-color: #f5f5f5;
            padding: 10px 15px;
            cursor: pointer;
            font-weight: bold;
            border-bottom: 1px solid #ddd;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .group-content {
            padding: 0;
        }
        .endpoint {
            border: 2px solid #aaa;
            border-radius: 4px;
            padding: 15px;
            position: relative;
            margin: 15px;
            background-color: #fff;
        }
        .endpoint:last-child {
            margin-bottom: 15px;
        }
        .endpoint-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 10px;
        }
        .endpoint-name {
            font-weight: bold;
            font-size: 1.1em;
        }
        .endpoint-description {
            color: #666;
            margin-bottom: 10px;
        }
        .method {
            display: inline-block;
            padding: 4px 8px;
            border-radius: 3px;
            font-weight: bold;
            color: white;
            background-color: #999;
            margin-right: 10px;
        }
        .method.get {
            background-color: #61affe;
        }
        .method.post {
            background-color: #49cc90;
        }
        .method.put {
            background-color: #fca130;
        }
        .method.delete {
            background-color: #f93e3e;
        }
        .method.unknown {
            background-color: #999;
        }
        .url {
            font-family: monospace;
            padding: 4px 8px;
            background-color: #f5f5f5;
            border-radius: 3px;
        }
        .parameters {
            margin-top: 15px;
        }
        .parameter-group {
            margin-bottom: 15px;
        }
        .parameter-group-title {
            font-weight: bold;
            margin-bottom: 5px;
            color: #555;
        }
        .parameter-table {
            width: 100%;
            border-collapse: collapse;
        }
        .parameter-table th, .parameter-table td {
            text-align: left;
            padding: 8px;
            border-bottom: 1px solid #eee;
        }
        .parameter-table th {
            background-color: #f5f5f5;
            font-weight: normal;
            color: #666;
        }
        .parameter-table input {
            width: 100%;
            padding: 6px;
            border: 1px solid #ddd;
            border-radius: 3px;
            font-family: monospace;
        }
        .labels {
            display: flex;
            gap: 5px;
            margin-top: 5px;
        }
        .label {
            font-size: 0.8em;
            padding: 2px 6px;
            border-radius: 10px;
            background-color: #eee;
        }
        .label.deprecated {
            background-color: #f93e3e;
            color: white;
        }
        .label.production {
            background-color: #49cc90;
            color: white;
        }
        .label.development {
            background-color: #fca130;
            color: white;
        }
        .toggle-btn {
            background: none;
            border: none;
            font-size: 1.2em;
            color: #666;
            pointer-events: none; /* Make the button not capture clicks */
        }
        .hidden {
            display: none;
        }
        .try-it-btn {
            background-color: #4CAF50;
            color: white;
            border: none;
            padding: 8px 16px;
            border-radius: 4px;
            cursor: pointer;
            font-weight: bold;
            margin-top: 15px;
            width: 100%;
        }
        .try-it-btn:hover {
            background-color: #45a049;
        }
        .try-it-section {
            margin-top: 20px;
            border-top: 1px solid #eee;
            padding-top: 20px;
        }
        .response-section {
            margin-top: 20px;
            border: 1px solid #ddd;
            border-radius: 4px;
            overflow: hidden;
        }
        .response-header {
            background-color: #f5f5f5;
            padding: 10px 15px;
            font-weight: bold;
            border-bottom: 1px solid #ddd;
            display: flex;
            justify-content: space-between;
        }
        .response-status {
            padding: 4px 8px;
            border-radius: 3px;
            font-weight: bold;
            color: white;
        }
        .response-status.success {
            background-color: #49cc90;
        }
        .response-status.error {
            background-color: #f93e3e;
        }
        .response-content {
            padding: 15px;
            background-color: #f8f8f8;
        }
        .response-headers-title, .response-body-title {
            font-weight: bold;
            margin-bottom: 8px;
            color: #555;
        }
        .response-headers {
            font-family: monospace;
            white-space: pre-wrap;
            margin-bottom: 15px;
            padding: 10px;
            background-color: #f0f0f0;
            border-radius: 3px;
            border-left: 3px solid #ddd;
        }
        .response-body {
            font-family: monospace;
            white-space: pre-wrap;
            padding: 10px;
            background-color: #f0f0f0;
            border-radius: 3px;
            border-left: 3px solid #ddd;
        }
        .response-divider {
            height: 1px;
            background-color: #eee;
            margin: 15px 0;
        }
        .textarea-wrapper {
            margin-top: 10px;
        }
        .textarea-wrapper textarea {
            width: 100%;
            min-height: 100px;
            padding: 8px;
            border: 1px solid #ddd;
            border-radius: 3px;
            font-family: monospace;
            resize: vertical;
        }
        .loading {
            text-align: center;
            padding: 40px;
            font-size: 1.2em;
            color: #666;
        }
        .error-message {
            background-color: #f8d7da;
            color: #721c24;
            padding: 15px;
            border-radius: 4px;
            margin-bottom: 20px;
            border: 1px solid #f5c6cb;
        }
        /* Custom header highlighting */
        .header-highlight {
            display: inline-block;
            padding: 2px 4px;
            border-radius: 3px;
            color: white;
            font-weight: bold;
        }
        .header-highlight.green {
            background-color: #49cc90;
        }
        .header-highlight.red {
            background-color: #f93e3e;
        }
        .header-highlight.blue {
            background-color: #61affe;
        }
        .header-highlight.orange {
            background-color: #fca130;
        }
        .header-highlight.purple {
            background-color: #9012fe;
        }
        .header-highlight.black {
            background-color: #333;
        }
    </style>
</head>
<body>
    <header>
        <h1 id="api-title"></h1>
        <p id="build">Build: <span id="build-number"></span>, started @ <span id="started-time"></span></p>
    </header>
    
    <!-- Global Auth Headers Section -->
    <div class="auth-section">
        <div class="auth-title">
            <span>Authentication Headers</span>
            <span id="auth-saved-indicator" class="auth-saved-indicator">Saved</span>
            <span id="auth-default" class="auth-default"></span>
        </div>
        <table class="auth-table" id="global-auth-table">
       
            <tbody>
                <!-- Auth headers will be added here dynamically -->
            </tbody>
        </table>

    </div>
    
    <div id="api-docs">
        <div class="loading">Loading API documentation...</div>
    </div>

    <script>
        // Cookie helper functions
        function setCookie(name, value, days) {
            const date = new Date();
            date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
            const expires = "expires=" + date.toUTCString();
            document.cookie = name + "=" + encodeURIComponent(value) + ";" + expires + ";path=/";
        }

        function getCookie(name) {
            const nameEQ = name + "=";
            const ca = document.cookie.split(';');
            for (let i = 0; i < ca.length; i++) {
                let c = ca[i];
                while (c.charAt(0) === ' ') c = c.substring(1, c.length);
                if (c.indexOf(nameEQ) === 0) return decodeURIComponent(c.substring(nameEQ.length, c.length));
            }
            return null;
        }

        // Save accordion state to cookie
        function saveAccordionState(groupId, isOpen) {
            setCookie("accordion_" + groupId, isOpen ? '1' : '0', 30); // Save for 30 days
        }

        // Show saved indicator briefly
        function showSavedIndicator() {
            const indicator = document.getElementById('auth-saved-indicator');
            indicator.style.display = 'inline';
            indicator.style.animation = 'none';
            // Trigger reflow
            void indicator.offsetWidth;
            indicator.style.animation = 'fadeOut 2s forwards';
        }

        // Save a specific auth value to cookie
        function saveAuthToCookie(name, value) {
            if (value) {
                setCookie("auth_" + name, value, 30); // Save for 30 days
                showSavedIndicator();
            }
        }

        // Function to highlight response headers based on configuration
        function highlightResponseHeaders(headerText, responseHeaders) {
            if (!responseHeaders || !responseHeaders.length) {
                return headerText;
            }

            let highlightedText = headerText;
            
            // Create a map for faster lookup
            const headerMap = {};
            responseHeaders.forEach(function(header) {
                if (!headerMap[header.name]) {
                    headerMap[header.name] = [];
                }
                headerMap[header.name].push({
                    value: header.value,
                    color: header.color || 'blue'
                });
            });

            // Process each line of the header text
            const lines = highlightedText.split('\n');
            const processedLines = lines.map(function(line) {
                const colonIndex = line.indexOf(':');
                if (colonIndex > 0) {
                    const headerName = line.substring(0, colonIndex).trim();
                    const headerValue = line.substring(colonIndex + 1).trim();
                    
                    if (headerMap[headerName]) {
                        // Check if this header value should be highlighted
                        const matchingRule = headerMap[headerName].find(function(rule) {
                            return rule.value === headerValue || rule.value === '*';
                        });
                        
                        if (matchingRule) {
                            return headerName + ": <span class=\"header-highlight " + matchingRule.color + "\">" + headerValue + "</span>";
                        }
                    }
                }
                return line;
            });

            return processedLines.join('\n');
        }

        // Function to initialize the API documentation
        function initializeApiDocs(apiDoc) {
            // Clear loading message
            const apiDocsContainer = document.getElementById('api-docs');
            apiDocsContainer.innerHTML = '';
            
            // Set the document title
            document.title = apiDoc.title || 'API Documentation';
            document.getElementById('api-title').textContent = apiDoc.title || 'API Documentation';

            // Set build
            document.getElementById('build-number').textContent = apiDoc.buildTag || 'v0.0.0';

            // Set service started time
            document.getElementById('started-time').textContent = apiDoc.started || 'yyyy-mm-dd hh:mm:ss';

            document.getElementById('auth-default').textContent = apiDoc.authDefaultMode ? 'Auth default: on' : 'Auth default: off';
            
            // Set header color if specified
            if (apiDoc.headerColor) {
                document.querySelector('header').style.backgroundColor = apiDoc.headerColor;
            }

            // Collect all unique auth headers from all endpoints
            const uniqueAuthHeaders = [];
            const authHeaderMap = new Map();

            apiDoc.groups.forEach(function(group) {
                group.endpoints.forEach(function(endpoint) {
                    if (endpoint.authHeaderOn && endpoint.authHeaders && endpoint.authHeaders.length > 0) {
                        endpoint.authHeaders.forEach(function(header) {
                            const key = header.name;
                            if (!authHeaderMap.has(key)) {
                                authHeaderMap.set(key, header);
                                uniqueAuthHeaders.push(header);
                            }
                        });
                    }
                });
            });

            // Populate global auth headers table
            const globalAuthTable = document.getElementById('global-auth-table').getElementsByTagName('tbody')[0];
            globalAuthTable.innerHTML = ''; // Clear existing rows
            
            uniqueAuthHeaders.forEach(function(header) {
                const row = globalAuthTable.insertRow();
                row.innerHTML = '<td>' + header.name + '</td>' +
                    '<td><input type="text" id="global-auth-' + header.name + '" placeholder="Enter ' + header.name + '"></td>';
            });

            // Add blur event listeners to auth input fields
            uniqueAuthHeaders.forEach(function(header) {
                const input = document.getElementById("global-auth-" + header.name);
                if (input) {
                    input.addEventListener('blur', function() {
                        saveAuthToCookie(header.name, this.value);
                    });
                }
            });

            // Generate the API documentation HTML
            apiDoc.groups.forEach(function(group, groupIndex) {
                const groupElement = document.createElement('div');
                groupElement.className = 'group';
                
                const groupHeader = document.createElement('div');
                groupHeader.className = 'group-header';
                groupHeader.innerHTML = '<span>' + group.name + '</span>' +
                    '<button class="toggle-btn" data-target="group-' + groupIndex + '">+</button>';
                groupElement.appendChild(groupHeader);
                
                const groupContent = document.createElement('div');
                groupContent.className = 'group-content hidden'; // Start with hidden class
                groupContent.id = 'group-' + groupIndex;
                
                group.endpoints.forEach(function(endpoint, endpointIndex) {
                    const endpointElement = document.createElement('div');
                    endpointElement.className = 'endpoint';
                    endpointElement.id = 'endpoint-' + groupIndex + '-' + endpointIndex;
                    
                    // Method class for styling
                    const methodClass = endpoint.method ? endpoint.method.toLowerCase() : 'unknown';
                    const isExecutable = endpoint.method && endpoint.method !== 'UNKNOWN';
                    
                    // Create endpoint header
                    const endpointHeader = document.createElement('div');
                    endpointHeader.className = 'endpoint-header';
                    endpointHeader.innerHTML = '<div>' +
                        '<span class="method ' + methodClass + '">' + (endpoint.method || 'UNKNOWN') + '</span>' +
                        '<span class="endpoint-name">' + endpoint.name + '</span>' +
                        '</div>';
                    
                    // Add labels if they exist
                    if (endpoint.label && endpoint.label.length > 0) {
                        const labelsDiv = document.createElement('div');
                        labelsDiv.className = 'labels';
                        
                        endpoint.label.forEach(function(label) {
                            const labelSpan = document.createElement('span');
                            labelSpan.className = 'label ' + label.toLowerCase();
                            labelSpan.textContent = label;
                            labelsDiv.appendChild(labelSpan);
                        });
                        
                        endpointHeader.appendChild(labelsDiv);
                    }
                    
                    endpointElement.appendChild(endpointHeader);
                    
                    // Add description
                    if (endpoint.description) {
                        const descriptionDiv = document.createElement('div');
                        descriptionDiv.className = 'endpoint-description';
                        descriptionDiv.textContent = endpoint.description;
                        endpointElement.appendChild(descriptionDiv);
                    }
                    
                    // Add URL
                    const urlDiv = document.createElement('div');
                    urlDiv.innerHTML = '<span class="url">' + endpoint.url + '</span>';
                    endpointElement.appendChild(urlDiv);
                    
                    // Parameters section
                    const parametersDiv = document.createElement('div');
                    parametersDiv.className = 'parameters';
                    
                    // Header Parameters (excluding auth headers which are now global)
                    if (endpoint.headerParameters && endpoint.headerParameters.length > 0) {
                        parametersDiv.appendChild(createParameterTable('Header Parameters', endpoint.headerParameters, isExecutable));
                    }
                    
                    // Query Parameters
                    if (endpoint.queryParameters && endpoint.queryParameters.length > 0) {
                        parametersDiv.appendChild(createParameterTable('Query Parameters', endpoint.queryParameters, isExecutable));
                    }
                    
                    // Path Parameters
                    if (endpoint.pathParameters && endpoint.pathParameters.length > 0) {
                        parametersDiv.appendChild(createParameterTable('Path Parameters', endpoint.pathParameters, isExecutable));
                    }
                    
                    // Body Parameters
                    if (endpoint.bodyParameters && endpoint.bodyParameters.length > 0) {
                        const bodyParamGroup = document.createElement('div');
                        bodyParamGroup.className = 'parameter-group';
                        
                        const bodyParamTitle = document.createElement('div');
                        bodyParamTitle.className = 'parameter-group-title';
                        // Create a string with all body parameter names and types
                        const paramInfo = endpoint.bodyParameters.map(function(param) {
                            return param.name + ' (' + param.type + ')';
                        }).join(', ');
                        bodyParamTitle.textContent = 'Body Parameters: ' + paramInfo;
                        bodyParamGroup.appendChild(bodyParamTitle);
                        
                        if (isExecutable) {
                            const textareaWrapper = document.createElement('div');
                            textareaWrapper.className = 'textarea-wrapper';
                            
                            // Use example JSON if available, otherwise create a default structure
                            let defaultJsonText = '';
                            
                            // Check if we have body parameters with types that match examples in exampleJsons
                            if (apiDoc.exampleJsons) {
                                // Find the first body parameter with a type that matches an example
                                const paramWithExample = endpoint.bodyParameters.find(function(param) {
                                    return apiDoc.exampleJsons[param.type];
                                });
                                
                                if (paramWithExample) {
                                    // Use the example JSON for this parameter type
                                    defaultJsonText = apiDoc.exampleJsons[paramWithExample.type];
                                    
                                    // Try to parse it to ensure it's valid JSON
                                    try {
                                        // If it's already a JSON object, stringify it properly
                                        const parsedJson = JSON.parse(defaultJsonText);
                                        defaultJsonText = JSON.stringify(parsedJson, null, 2);
                                    } catch (e) {
                                        // If it's not valid JSON, use it as is (it might be a string representation)
                                        console.warn("Example JSON for " + paramWithExample.type + " is not valid JSON");
                                    }
                                }
                            }
                            
                            // If no matching example was found, create a default structure
                            if (!defaultJsonText) {
                                let defaultJson = {};
                                endpoint.bodyParameters.forEach(function(param) {
                                    // Handle array types
                                    if (param.type.startsWith('[]')) {
                                        defaultJson[param.name] = [];
                                    } else {
                                        defaultJson[param.name] = '';
                                    }
                                });
                                defaultJsonText = JSON.stringify(defaultJson, null, 2);
                            }
                            
                            const textareaHtml = '<textarea id="body-param-' + groupIndex + '-' + endpointIndex + 
                                '" placeholder="Enter request body">' + defaultJsonText + '</textarea>';
                            textareaWrapper.innerHTML = textareaHtml;
                            bodyParamGroup.appendChild(textareaWrapper);
                        } else {
                            const table = document.createElement('table');
                            table.className = 'parameter-table';
                            
                            // Table header
                            const thead = document.createElement('thead');
                            thead.innerHTML = '<tr>' +
                                '<th>Name</th>' +
                                '<th>Type</th>' +
                                '</tr>';
                            table.appendChild(thead);
                            
                            // Table body
                            const tbody = document.createElement('tbody');
                            endpoint.bodyParameters.forEach(function(param) {
                                const tr = document.createElement('tr');
                                tr.innerHTML = '<td>' + param.name + '</td>' +
                                    '<td>' + param.type + '</td>';
                                tbody.appendChild(tr);
                            });
                            table.appendChild(tbody);
                            
                            bodyParamGroup.appendChild(table);
                        }
                        
                        parametersDiv.appendChild(bodyParamGroup);
                    }
                    
                    endpointElement.appendChild(parametersDiv);
                    
                    // Add "Execute" button for executable endpoints
                    if (isExecutable) {
                        const tryItButton = document.createElement('button');
                        tryItButton.className = 'try-it-btn';
                        tryItButton.textContent = 'Execute';
                        tryItButton.setAttribute('data-endpoint', groupIndex + '-' + endpointIndex);
                        tryItButton.addEventListener('click', executeRequest);
                        endpointElement.appendChild(tryItButton);
                        
                        // Add response section (initially hidden)
                        const responseSection = document.createElement('div');
                        responseSection.className = 'response-section hidden';
                        responseSection.id = 'response-' + groupIndex + '-' + endpointIndex;
                        
                        const responseHtml = '<div class="response-header">' +
                            '<span>Response</span>' +
                            '<span class="response-status" id="status-' + groupIndex + '-' + endpointIndex + '"></span>' +
                            '</div>' +
                            '<div class="response-content">' +
                            '<div class="response-headers-title">Headers:</div>' +
                            '<div class="response-headers" id="response-headers-' + groupIndex + '-' + endpointIndex + '"></div>' +
                            '<div class="response-body-title">Body:</div>' +
                            '<div class="response-body" id="response-body-' + groupIndex + '-' + endpointIndex + '"></div>' +
                            '</div>';
                        
                        responseSection.innerHTML = responseHtml;
                        endpointElement.appendChild(responseSection);
                    }
                    
                    groupContent.appendChild(endpointElement);
                });
                
                groupElement.appendChild(groupContent);
                apiDocsContainer.appendChild(groupElement);
            });
            
            // Add toggle functionality to the entire group header
            document.querySelectorAll('.group-header').forEach(function(header) {
                header.addEventListener('click', function() {
                    const targetId = this.querySelector('.toggle-btn').getAttribute('data-target');
                    const targetElement = document.getElementById(targetId);
                    const toggleButton = this.querySelector('.toggle-btn');
                    
                    if (targetElement.classList.contains('hidden')) {
                        targetElement.classList.remove('hidden');
                        toggleButton.textContent = '-';
                        saveAccordionState(targetId, true);
                    } else {
                        targetElement.classList.add('hidden');
                        toggleButton.textContent = '+';
                        saveAccordionState(targetId, false);
                    }
                });
            });
            
            // Load accordion states from cookies
            loadAccordionStates();
            
            // Load auth values from cookies
            loadAuthFromCookies(uniqueAuthHeaders);
        }
        
        // Function to create parameter tables
        function createParameterTable(title, parameters, isExecutable) {
            const parameterGroup = document.createElement('div');
            parameterGroup.className = 'parameter-group';
            
            const parameterTitle = document.createElement('div');
            parameterTitle.className = 'parameter-group-title';
            parameterTitle.textContent = title;
            parameterGroup.appendChild(parameterTitle);
            
            const table = document.createElement('table');
            table.className = 'parameter-table';
            
            // Table header
            const thead = document.createElement('thead');
            thead.innerHTML = '<tr>' +
                '<th>Name</th>' +
                '<th>Type</th>' +
                (isExecutable ? '<th>Value</th>' : '') +
                '</tr>';
            table.appendChild(thead);
            
            // Table body
            const tbody = document.createElement('tbody');
            parameters.forEach(function(param, index) {
                const tr = document.createElement('tr');
                tr.innerHTML = '<td>' + param.name + '</td>' +
                    '<td>' + param.type + '</td>' +
                    (isExecutable ? '<td><input type="text" id="' + title.toLowerCase().replace(' ', '-') + '-' + param.name + 
                    '" placeholder="Enter ' + param.name + '"></td>' : '');
                tbody.appendChild(tr);
            });
            table.appendChild(tbody);
            
            parameterGroup.appendChild(table);
            return parameterGroup;
        }
        
        // Load accordion states from cookies
        function loadAccordionStates() {
            document.querySelectorAll('.group-content').forEach(function(content) {
                const groupId = content.id;
                const isOpen = getCookie("accordion_" + groupId) === '1';
                
                if (isOpen) {
                    content.classList.remove('hidden');
                    const button = document.querySelector('button[data-target="' + groupId + '"]');
                    if (button) {
                        button.textContent = '-';
                    }
                }
            });
        }
        
        // Load auth values from cookies
        function loadAuthFromCookies(uniqueAuthHeaders) {
            uniqueAuthHeaders.forEach(function(header) {
                const cookieValue = getCookie("auth_" + header.name);
                if (cookieValue) {
                    const input = document.getElementById("global-auth-" + header.name);
                    if (input) {
                        input.value = cookieValue;
                    }
                }
            });
        }
        
        // Function to execute request
        function executeRequest(e) {
            const endpointId = e.target.getAttribute('data-endpoint');
            const idParts = endpointId.split('-');
            
            const groupIndex = parseInt(idParts[0]);
            const endpointIndex = parseInt(idParts[1]);
            
            // Get the current apiDoc from the window object
            const apiDoc = window.apiDocData;
            const endpoint = apiDoc.groups[groupIndex].endpoints[endpointIndex];
            
            // Show response section
            const responseSection = document.getElementById('response-' + groupIndex + '-' + endpointIndex);
            responseSection.classList.remove('hidden');
            
            // Build request URL
            let url = endpoint.url;
            
            // Add path parameters
            if (endpoint.pathParameters && endpoint.pathParameters.length > 0) {
                endpoint.pathParameters.forEach(function(param) {
                    const value = document.getElementById('path-parameters-' + param.name).value;
                    if (value) {
                        // Replace placeholder in URL with actual value
                        // This is a simplification - in a real app, you'd need to handle the actual URL pattern
                        url = url.replace('{' + param.name + '}', value);
                    }
                });
            }
            
            // Add query parameters
            if (endpoint.queryParameters && endpoint.queryParameters.length > 0) {
                const queryParams = [];
                endpoint.queryParameters.forEach(function(param) {
                    const value = document.getElementById('query-parameters-' + param.name).value;
                    if (value) {
                        queryParams.push(param.name + '=' + encodeURIComponent(value));
                    }
                });
                
                if (queryParams.length > 0) {
                    url += (url.includes('?') ? '&' : '?') + queryParams.join('&');
                }
            }
            
            // Build request headers
            const headers = {};
            
            // Get all unique auth headers
            const uniqueAuthHeaders = [];
            const authHeaderMap = new Map();
            
            apiDoc.groups.forEach(function(group) {
                group.endpoints.forEach(function(endpoint) {
                    if (endpoint.authHeaderOn && endpoint.authHeaders && endpoint.authHeaders.length > 0) {
                        endpoint.authHeaders.forEach(function(header) {
                            const key = header.name;
                            if (!authHeaderMap.has(key)) {
                                authHeaderMap.set(key, header);
                                uniqueAuthHeaders.push(header);
                            }
                        });
                    }
                });
            });
            
            // Add global auth headers to all requests
            uniqueAuthHeaders.forEach(function(header) {
                const value = document.getElementById('global-auth-' + header.name).value;
                if (value) {
                    headers[header.name] = value;
                }
            });
            
            // Add header parameters
            if (endpoint.headerParameters && endpoint.headerParameters.length > 0) {
                endpoint.headerParameters.forEach(function(param) {
                    const value = document.getElementById('header-parameters-' + param.name).value;
                    if (value) {
                        headers[param.name] = value;
                    }
                });
            }
            
            // Build request body
            let body = null;
            if (endpoint.bodyParameters && endpoint.bodyParameters.length > 0) {
                const bodyTextarea = document.getElementById('body-param-' + groupIndex + '-' + endpointIndex);
                if (bodyTextarea && bodyTextarea.value) {
                    try {
                        body = JSON.parse(bodyTextarea.value);
                        headers['Content-Type'] = 'application/json';
                    } catch (e) {
                        // Show error if JSON is invalid
                        document.getElementById('response-body-' + groupIndex + '-' + endpointIndex).textContent = 'Error: Invalid JSON in request body';
                        document.getElementById('status-' + groupIndex + '-' + endpointIndex).textContent = 'Error';
                        document.getElementById('status-' + groupIndex + '-' + endpointIndex).className = 'response-status error';
                        return;
                    }
                }
            }
            
            // Execute the request
            fetch(url, {
                method: endpoint.method,
                headers: headers,
                body: body ? JSON.stringify(body) : null
            })
            .then(function(response) {
                // Display status
                const statusElement = document.getElementById('status-' + groupIndex + '-' + endpointIndex);
                statusElement.textContent = response.status + ' ' + response.statusText;
                statusElement.className = 'response-status ' + (response.ok ? 'success' : 'error');
                
                // Display headers
                const responseHeaders = document.getElementById('response-headers-' + groupIndex + '-' + endpointIndex);
                let headersText = '';
                response.headers.forEach(function(value, name) {
                    headersText += name + ': ' + value + '\n';
                });
                
                // Apply header highlighting if configured
                if (apiDoc.responseHeaders && apiDoc.responseHeaders.length > 0) {
                    responseHeaders.innerHTML = highlightResponseHeaders(headersText, apiDoc.responseHeaders);
                } else {
                    responseHeaders.textContent = headersText;
                }
                
                // Check content type for JSON
                const contentType = response.headers.get('content-type');
                if (contentType && contentType.includes('application/json')) {
                    return response.json().then(function(data) {
                        return {
                            data: JSON.stringify(data, null, 2),
                            isJson: true
                        };
                    });
                } else {
                    return response.text().then(function(text) {
                        return {
                            data: text,
                            isJson: false
                        };
                    });
                }
            })
            .then(function(result) {
                // Display response body
                const responseBody = document.getElementById('response-body-' + groupIndex + '-' + endpointIndex);
                responseBody.textContent = result.data;
                
                // Add syntax highlighting for JSON (simplified version)
                if (result.isJson) {
                    responseBody.classList.add('json');
                }
            })
            .catch(function(error) {
                // Display error
                document.getElementById('response-body-' + groupIndex + '-' + endpointIndex).textContent = 'Error: ' + error.message;
                document.getElementById('status-' + groupIndex + '-' + endpointIndex).textContent = 'Error';
                document.getElementById('status-' + groupIndex + '-' + endpointIndex).className = 'response-status error';
            });
        }
        
        // Fetch the API documentation JSON from the server
        fetch('/microapidocs/doc.json')
            .then(function(response) {
                if (!response.ok) {
                    throw new Error('Failed to load API documentation: ' + response.status + ' ' + response.statusText);
                }
                return response.json();
            })
            .then(function(data) {
                // Store the API data globally for access in event handlers
                window.apiDocData = data;
                
                // Initialize the documentation with the fetched data
                initializeApiDocs(data);
            })
            .catch(function(error) {
                // Display error message
                const apiDocsContainer = document.getElementById('api-docs');
                apiDocsContainer.innerHTML = '<div class="error-message">' +
                    '<strong>Error loading API documentation:</strong> ' + error.message +
                    '</div>';
                console.error('Error loading API documentation:', error);
            });
    </script>
</body>
</html>
`

	ctx.Data(200, "text/html; charset=utf-8", []byte(indexHtml))
	return

}
