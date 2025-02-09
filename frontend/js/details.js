const urlParams = new URLSearchParams(window.location.search);
const campaignId = urlParams.get("id");

async function loadCampaignDetails() {
    try {
        const response = await fetch(`/campaigns/${campaignId}`); // Adjust API URL
        const campaign = await response.json();

        document.getElementById("campaign-title").innerText = campaign.title;
        document.getElementById("campaign-description").innerText = campaign.description;
        document.getElementById("campaign-goal").innerText = campaign.target_amount;
        document.getElementById("campaign-raised").innerText = campaign.amount_raised;

        const mediaContainer = document.getElementById("campaign-media");
        mediaContainer.innerHTML = `<img src="${campaign.media_path}" alt="Campaign Image" class="campaign-img">`;


        document.getElementById("donate-btn").addEventListener("click", () => {
            window.location.href = `/static/donation.html?id=${campaignId}`;
        });

    } catch (error) {
        console.error("Error loading campaign:", error);
    }
}

loadCampaignDetails();
