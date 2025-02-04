document.addEventListener("DOMContentLoaded", async () => {
    const campaignList = document.getElementById("campaign-list");
    const searchForm = document.getElementById("searchForm");
    const searchInput = document.getElementById("searchInput");

    // Check if the elements exist
   if (!searchForm || !searchInput || !campaignList) {
        console.error("Required DOM elements are missing.");
        return;
    }

    // Function to fetch campaigns
    async function fetchCampaigns(query = "") {
        const url = query ? `http://localhost:8080/campaigns/search?query=${query}` : "http://localhost:8080/campaigns/";

        try {
            const response = await fetch(url);

            // Log the raw response to check if it's valid
            console.log("Response:", response);

            if (!response.ok) {
                console.error('Failed to fetch campaigns', response.status);
                return;
            }

            let campaigns = await response.json();
            if (campaigns === null) {
                campaigns = [];
            }
            // Log the parsed campaigns to ensure they are in the expected format
            console.log("Campaigns:", campaigns);

            // Check if campaigns is an array
            if (!Array.isArray(campaigns)) {
                console.error("Invalid response format: expected an array of campaigns.");
                return;
            }

            campaignList.innerHTML = ''; // Clear the campaign list before displaying new results

            if (campaigns.length === 0) {
                campaignList.innerHTML = '<p>No campaigns found.</p>';
            } else {
                campaigns.forEach(campaign => {
                    const campaignDiv = document.createElement("div");
                    campaignDiv.className = "campaign";
                    campaignDiv.innerHTML = `
                        <h3>${campaign.title}</h3>
                        <p>${campaign.campaign_id}</p>
                        <p>${campaign.description}</p>
                        <p><strong>Goal:</strong> $${campaign.target_amount}</p>
                        <a href="/static/donation.html?id=${campaign.campaign_id}" class="btn">Donate</a>
                    `;
                    campaignList.appendChild(campaignDiv);
                });
            }
        } catch (error) {
            console.error("Error fetching campaigns:", error);
        }
    }

    // Initial fetch for all campaigns
    fetchCampaigns();

    // Event listener for the search form submission
    searchForm.addEventListener("submit", (e) => {
        e.preventDefault();
        const query = searchInput.value.trim();  // Get the query from the input field

        // Update the URL without reloading the page
       if (query) {
            history.pushState({}, "", `?search=${query}`); // Update URL
        } else {
            history.pushState({}, "", "/"); // Reset the URL when search is empty
        }

        fetchCampaigns(query);  // Fetch campaigns based on the search query
    });

    // Optionally handle back/forward navigation to maintain the state when user presses browser buttons
    window.addEventListener('popstate', () => {
        const queryParams = new URLSearchParams(window.location.search);
        const query = queryParams.get('search') || '';
        searchInput.value = query; // Optionally update the search input with the query
        fetchCampaigns(query); // Re-fetch campaigns based on the query in the URL
    });
});


function checkUserLoginStatus() {
    const token = localStorage.getItem("token");

    if (!token) {
        document.getElementById("campaign-menu").style.display = "none";
        return;
    }

    const decoded = jwt_decode(token);

    const currentTime = Date.now() / 1000; // Current time in seconds
    if (decoded.exp < currentTime) {
        document.getElementById("campaign-menu").style.display = "none";
    } else {
        document.getElementById("campaign-menu").style.display = "block";
    }
}

window.onload = checkUserLoginStatus;