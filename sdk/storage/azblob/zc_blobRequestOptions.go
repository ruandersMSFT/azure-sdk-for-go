package azblob

import (
	"net/url"
	"strings"
)

type DeleteBlobOptions struct {
	// Required if the blob has associated snapshots. Specify one of the following two options: include: Delete the base blob
	// and all of its snapshots. only: Delete only the blob's snapshots and not the blob itself
	DeleteSnapshots *DeleteSnapshotsOptionType

	LeaseAccessConditions    *LeaseAccessConditions
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *DeleteBlobOptions) pointers() (*BlobDeleteOptions, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}

	basics := BlobDeleteOptions{
		DeleteSnapshots: o.DeleteSnapshots,
	}

	return &basics, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type DownloadBlobOptions struct {
	// When set to true and specified together with the Range, the service returns the MD5 hash for the range, as long as the
	// range is less than or equal to 4 MB in size.
	RangeGetContentMd5 *bool

	// Optional, you can specify whether a particular range of the blob is read
	Offset *int64
	Count  *int64

	LeaseAccessConditions    *LeaseAccessConditions
	CpkInfo                  *CpkInfo
	CpkScopeInfo             *CpkScopeInfo
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *DownloadBlobOptions) pointers() (blobDownloadOptions *BlobDownloadOptions, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil
	}

	offset := int64(0)
	count := int64(CountToEnd)

	if o.Offset != nil {
		offset = *o.Offset
	}

	if o.Count != nil {
		count = *o.Count
	}

	basics := BlobDownloadOptions{
		RangeGetContentMd5: o.RangeGetContentMd5,
		RangeParameter: httpRange{
			offset: offset,
			count:  count,
		}.pointers(),
	}

	return &basics, o.LeaseAccessConditions, o.CpkInfo, o.ModifiedAccessConditions
}

type SetTierOptions struct {
	// Optional: Indicates the priority with which to rehydrate an archived blob.
	RehydratePriority *RehydratePriority

	LeaseAccessConditions    *LeaseAccessConditions
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *SetTierOptions) pointers() (blobSetTierOptions *BlobSetTierOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}

	basics := BlobSetTierOptions{RehydratePriority: o.RehydratePriority}
	return &basics, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type GetBlobPropertiesOptions struct {
	LeaseAccessConditions    *LeaseAccessConditions
	CpkInfo                  *CpkInfo
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *GetBlobPropertiesOptions) pointers() (blobGetPropertiesOptions *BlobGetPropertiesOptions, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil
	}

	return nil, o.LeaseAccessConditions, o.CpkInfo, o.ModifiedAccessConditions
}

type SetBlobHTTPHeadersOptions struct {
	LeaseAccessConditions    *LeaseAccessConditions
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *SetBlobHTTPHeadersOptions) pointers() (blobSetHttpHeadersOptions *BlobSetHTTPHeadersOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}

	return nil, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type SetBlobMetadataOptions struct {
	LeaseAccessConditions    *LeaseAccessConditions
	CpkInfo                  *CpkInfo
	CpkScopeInfo             *CpkScopeInfo
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *SetBlobMetadataOptions) pointers() (leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil
	}

	return o.LeaseAccessConditions, o.CpkInfo, o.CpkScopeInfo, o.ModifiedAccessConditions
}

type CreateBlobSnapshotOptions struct {
	Metadata                 *map[string]string
	LeaseAccessConditions    *LeaseAccessConditions
	CpkInfo                  *CpkInfo
	CpkScopeInfo             *CpkScopeInfo
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *CreateBlobSnapshotOptions) pointers() (blobSetMetadataOptions *BlobCreateSnapshotOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions, leaseAccessConditions *LeaseAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil, nil
	}

	basics := BlobCreateSnapshotOptions{
		Metadata: o.Metadata,
	}

	return &basics, o.CpkInfo, o.CpkScopeInfo, o.ModifiedAccessConditions, o.LeaseAccessConditions
}

type AcquireBlobLeaseOptions struct {
	// Specifies the duration of the lease, in seconds, or negative one (-1) for a lease that never expires. A non-infinite lease
	// can be between 15 and 60 seconds. A lease duration cannot be changed using renew or change.
	Duration int32

	// Proposed lease ID, in a GUID string format. The Blob service returns 400 (Invalid request) if the proposed lease ID is
	// not in the correct format. See Guid Constructor (String) for a list of valid GUID string formats.
	ProposedLeaseID *string

	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *AcquireBlobLeaseOptions) pointers() (blobAcquireLeaseOptions *BlobAcquireLeaseOptions, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil
	}

	basics := BlobAcquireLeaseOptions{
		Duration:        &o.Duration,
		ProposedLeaseId: o.ProposedLeaseID,
	}

	return &basics, o.ModifiedAccessConditions
}

type RenewBlobLeaseOptions struct {
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *RenewBlobLeaseOptions) pointers() (blobRenewLeaseOptions *BlobRenewLeaseOptions, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return nil, o.ModifiedAccessConditions
}

type ReleaseBlobLeaseOptions struct {
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *ReleaseBlobLeaseOptions) pointers() (blobReleaseLeaseOptions *BlobReleaseLeaseOptions, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return nil, o.ModifiedAccessConditions
}

type BreakBlobLeaseOptions struct {
	// For a break operation, proposed duration the lease should continue before it is broken, in seconds, between 0 and 60. This
	// break period is only used if it is shorter than the time remaining on the lease. If longer, the time remaining on the lease
	// is used. A new lease will not be available before the break period has expired, but the lease may be held for longer than
	// the break period. If this header does not appear with a break operation, a fixed-duration lease breaks after the remaining
	// lease period elapses, and an infinite lease breaks immediately.
	BreakPeriod              *int32
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *BreakBlobLeaseOptions) pointers() (blobBreakLeaseOptions *BlobBreakLeaseOptions, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil
	}

	if o.BreakPeriod != nil {
		period := leasePeriodPointer(*o.BreakPeriod)
		basics := BlobBreakLeaseOptions{
			BreakPeriod: period,
		}
		return &basics, o.ModifiedAccessConditions
	}

	return nil, o.ModifiedAccessConditions
}

// LeaseBreakNaturally tells ContainerURL's or BlobClient's BreakLease method to break the lease using service semantics.
const LeaseBreakNaturally = -1

func leasePeriodPointer(period int32) (p *int32) {
	if period != LeaseBreakNaturally {
		p = &period
	}
	return nil
}

type ChangeBlobLeaseOptions struct {
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *ChangeBlobLeaseOptions) pointers() (blobChangeLeaseOptions *BlobChangeLeaseOptions, modifiedAccessConditions *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return nil, o.ModifiedAccessConditions
}

type StartCopyBlobOptions struct {
	// Optional. Used to set blob tags in various blob operations.
	BlobTagsMap *map[string]string
	// Optional. Specifies a user-defined name-value pair associated with the blob. If no name-value pairs are specified, the
	// operation will copy the metadata from the source blob or file to the destination blob. If one or more name-value pairs
	// are specified, the destination blob is created with the specified metadata, and metadata is not copied from the source
	// blob or file. Note that beginning with version 2009-09-19, metadata names must adhere to the naming rules for C# identifiers.
	// See Naming and Referencing Containers, Blobs, and Metadata for more information.
	Metadata *map[string]string
	// Optional: Indicates the priority with which to rehydrate an archived blob.
	RehydratePriority *RehydratePriority
	// Overrides the sealed state of the destination blob. Service version 2019-12-12 and newer.
	SealBlob *bool
	// Optional. Indicates the tier to be set on the blob.
	Tier *AccessTier

	SourceModifiedAccessConditions *SourceModifiedAccessConditions
	ModifiedAccessConditions       *ModifiedAccessConditions
	LeaseAccessConditions          *LeaseAccessConditions
}

func (o *StartCopyBlobOptions) pointers() (blobStartCopyFromUrlOptions *BlobStartCopyFromURLOptions, sourceModifiedAccessConditions *SourceModifiedAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, leaseAccessConditions *LeaseAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil
	}

	basics := BlobStartCopyFromURLOptions{
		BlobTagsString:    SerializeBlobTagsToStrPtr(o.BlobTagsMap),
		Metadata:          o.Metadata,
		RehydratePriority: o.RehydratePriority,
		SealBlob:          o.SealBlob,
		Tier:              o.Tier,
	}

	return &basics, o.SourceModifiedAccessConditions, o.ModifiedAccessConditions, o.LeaseAccessConditions
}

type AbortCopyBlobOptions struct {
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *AbortCopyBlobOptions) pointers() (blobAbortCopyFromUrlOptions *BlobAbortCopyFromURLOptions, leaseAccessConditions *LeaseAccessConditions) {
	if o == nil {
		return nil, nil
	}
	return nil, o.LeaseAccessConditions
}

func SerializeBlobTagsToStrPtr(blobTagsMap *map[string]string) *string {
	if blobTagsMap == nil {
		return nil
	}
	tags := make([]string, 0)
	for key, val := range *blobTagsMap {
		tags = append(tags, url.QueryEscape(key)+"="+url.QueryEscape(val))
	}
	//tags = tags[:len(tags)-1]
	blobTagsString := strings.Join(tags, "&")
	return &blobTagsString
}

func SerializeBlobTags(blobTagsMap *map[string]string) *BlobTags {
	if blobTagsMap == nil {
		return nil
	}
	blobTagSet := make([]BlobTag, 0)
	for key, val := range *blobTagsMap {
		newKey, newVal := key, val
		blobTagSet = append(blobTagSet, BlobTag{Key: &newKey, Value: &newVal})
	}
	return &BlobTags{BlobTagSet: &blobTagSet}
}

type SetTagsBlobOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when storage analytics logging is enabled.
	RequestId *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
	// The version id parameter is an opaque DateTime value that, when present,
	// specifies the version of the blob to operate on. It's for service version 2019-10-10 and newer.
	VersionId *string
	// Optional header, Specifies the transactional crc64 for the body, to be validated by the service.
	TransactionalContentCrc64 *[]byte
	// Optional header, Specifies the transactional md5 for the body, to be validated by the service.
	TransactionalContentMd5 *[]byte

	BlobTagsMap *map[string]string

	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *SetTagsBlobOptions) pointers() (*BlobSetTagsOptions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil
	}

	options := &BlobSetTagsOptions{
		RequestId:                 o.RequestId,
		Tags:                      SerializeBlobTags(o.BlobTagsMap),
		Timeout:                   o.Timeout,
		TransactionalContentMd5:   o.TransactionalContentMd5,
		TransactionalContentCrc64: o.TransactionalContentCrc64,
		VersionId:                 o.VersionId,
	}

	return options, o.ModifiedAccessConditions
}

type GetTagsBlobOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when storage analytics logging is enabled.
	RequestId *string
	// The snapshot parameter is an opaque DateTime value that, when present, specifies the blob snapshot to retrieve.
	Snapshot *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
	// The version id parameter is an opaque DateTime value that, when present, specifies the version of the blob to operate on.
	// It's for service version 2019-10-10 and newer.
	VersionId *string

	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *GetTagsBlobOptions) pointers() (*BlobGetTagsOptions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil
	}

	options := &BlobGetTagsOptions{
		RequestId: o.RequestId,
		Snapshot:  o.Snapshot,
		Timeout:   o.Timeout,
		VersionId: o.VersionId,
	}

	return options, o.ModifiedAccessConditions
}