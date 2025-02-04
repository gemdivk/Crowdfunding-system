async function fetchUserCampaigns() {
    const token = localStorage.getItem("token");
    const decoded = jwt_decode(token);
    const userId = decoded.UserID;

    if (!userId) {
        console.error("User ID not found");
        return;
    }

    const response = await fetch(`/campaigns/user/${userId}`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`
        }
    });
    const data = await response.json();
    const campaignList = document.getElementById("user-campaign-list");
    campaignList.innerHTML = "";

    if (!data || data.length === 0) {
        campaignList.innerHTML = "<p>You have no campaigns.</p>";
        return;
    }

    data.forEach(campaign => {
        const listItem = document.createElement("li");
        listItem.classList.add("campaign-item");
        listItem.innerHTML = `
            <h3>${campaign.title}</h3>
            <p>${campaign.description}</p>
              <p>${campaign.campaign_id}</p>
            <p><strong>Target:</strong> $${campaign.target_amount}</p>
            <p><strong>Raised:</strong> $${campaign.amount_raised}</p>
            <p><strong>Status:</strong> ${campaign.status}</p>
             <button onclick="editCampaign(${campaign.campaign_id})">Edit</button>
              <button onclick="deleteCampaign(${campaign.campaign_id})">Delete</button>
            <hr>
        `;
        campaignList.appendChild(listItem);
    });
}


function showCreateForm() {
    document.getElementById('create-form').style.display = 'block';
}

function hideCreateForm() {
    document.getElementById('create-form').style.display = 'none';
}

// Function to handle create campaign submission
function createCampaign() {
    const title = document.getElementById('create-title').value;
    const description = document.getElementById('create-description').value;
    const targetAmount = document.getElementById('create-target-amount').value;
    const status = document.getElementById('create-status').value;
    const token = localStorage.getItem("token");
    const decoded = jwt_decode(token);
    const userId = decoded.UserID;

    if (!userId) {
        console.error("User ID not found");
        return;
    }
    const data = {
        user_id: userId,          // Renamed to match the expected backend field name
        title: title,    // Same for other properties
        description: description,
        target_amount: parseFloat(targetAmount),
        status: status,
    };

    // Send a POST request to create the campaign
    fetch('/campaigns/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(data),
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Campaign created successfully!');
                hideCreateForm();  // Hide form after successful submission
                fetchUserCampaigns();   // Refresh the campaign list
            } else {
                alert('Failed to create campaign');
            }
        })
        .catch(error => {
            console.error('Error creating campaign:', error);
            alert('An error occurred while creating the campaign.');
        });
}
fetchUserCampaigns();
async function deleteCampaign(campaign_id) {
    if (!confirm("Are you sure you want to delete this campaign?")) return;

    const response = await fetch(`/campaigns/${campaign_id}`, {
        method: "DELETE",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token")}`
        }
    });

    if (response.ok) {
        alert("Campaign deleted!");
        fetchMyCampaigns();
    } else {
        console.error("Error deleting campaign", response.status);
    }
}
async function editCampaign(campaign_id) {
    const newTitle = prompt("Enter new title:");
    const newDescription = prompt("Enter new description:");
    const newTargetAmount = prompt("Enter new target amount:");
    const token = localStorage.getItem("token");
    const decoded = jwt_decode(token);
    const userId = decoded.UserID;
    const status = prompt("Enter desired new status")
    if (!newTitle || !newDescription || !newTargetAmount || !status) {
        alert("All fields are required!");
        return;
    }
    const targetAmountNum = parseFloat(newTargetAmount)
    const campaignData = {
        title: newTitle,
        description: newDescription,
        target_amount: targetAmountNum,
        status:status,
    };

//    console.log("Sending data:", JSON.stringify(campaignData));

    fetch(`/campaigns/${campaign_id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(campaignData),
    })
        .then(response => response.json())
        .then(campaignData => {
            if (campaignData.success) {
                alert("Campaign updated successfully!");
                fetchUserCampaigns();  // Refresh the campaign list
            } else {
                alert(`Failed to update campaign: ${data.message || 'Unknown error'}`);
            }
        })
        .catch(error => {
            console.error("Error updating campaign", error);
            alert("An error occurred while updating the campaign.");
        });
}