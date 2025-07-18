// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Outputs
{

    [OutputType]
    public sealed class WorkflowTaskCreateGoogleCalendarEventTaskParams
    {
        /// <summary>
        /// Emails of attendees
        /// </summary>
        public readonly ImmutableArray<string> Attendees;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? CanGuestsInviteOthers;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? CanGuestsModifyEvent;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? CanGuestsSeeOtherGuests;
        /// <summary>
        /// Sets the video conference type attached to the meeting. Value must be one of `eventHangout`, `eventNamedHangout`, `hangoutsMeet`, `addOn`.
        /// </summary>
        public readonly string? ConferenceSolutionKey;
        /// <summary>
        /// The days until meeting
        /// </summary>
        public readonly int DaysUntilMeeting;
        /// <summary>
        /// The event description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? ExcludeWeekends;
        /// <summary>
        /// Meeting duration in format like '1 hour', '30 minutes'
        /// </summary>
        public readonly string MeetingDuration;
        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        public readonly bool? PostToIncidentTimeline;
        public readonly ImmutableArray<Outputs.WorkflowTaskCreateGoogleCalendarEventTaskParamsPostToSlackChannel> PostToSlackChannels;
        /// <summary>
        /// Send an email to the attendees notifying them of the event. Value must be one of true or false
        /// </summary>
        public readonly bool? SendUpdates;
        /// <summary>
        /// The event summary
        /// </summary>
        public readonly string Summary;
        public readonly string? TaskType;
        /// <summary>
        /// Time of meeting in format HH:MM
        /// </summary>
        public readonly string TimeOfMeeting;
        /// <summary>
        /// A valid IANA time zone name.
        /// </summary>
        public readonly string? TimeZone;

        [OutputConstructor]
        private WorkflowTaskCreateGoogleCalendarEventTaskParams(
            ImmutableArray<string> attendees,

            bool? canGuestsInviteOthers,

            bool? canGuestsModifyEvent,

            bool? canGuestsSeeOtherGuests,

            string? conferenceSolutionKey,

            int daysUntilMeeting,

            string description,

            bool? excludeWeekends,

            string meetingDuration,

            bool? postToIncidentTimeline,

            ImmutableArray<Outputs.WorkflowTaskCreateGoogleCalendarEventTaskParamsPostToSlackChannel> postToSlackChannels,

            bool? sendUpdates,

            string summary,

            string? taskType,

            string timeOfMeeting,

            string? timeZone)
        {
            Attendees = attendees;
            CanGuestsInviteOthers = canGuestsInviteOthers;
            CanGuestsModifyEvent = canGuestsModifyEvent;
            CanGuestsSeeOtherGuests = canGuestsSeeOtherGuests;
            ConferenceSolutionKey = conferenceSolutionKey;
            DaysUntilMeeting = daysUntilMeeting;
            Description = description;
            ExcludeWeekends = excludeWeekends;
            MeetingDuration = meetingDuration;
            PostToIncidentTimeline = postToIncidentTimeline;
            PostToSlackChannels = postToSlackChannels;
            SendUpdates = sendUpdates;
            Summary = summary;
            TaskType = taskType;
            TimeOfMeeting = timeOfMeeting;
            TimeZone = timeZone;
        }
    }
}
