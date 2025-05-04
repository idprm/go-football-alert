SELECT DATE(created_at) as created_at, subject, status, COUNT(1) as total, SUM(amount) as revenue 
FROM fb_alert_test.transactions WHERE DATE(created_at) BETWEEN DATE('2025-04-26')
AND DATE('2025-04-26') GROUP BY DATE(created_at), subject, status 
ORDER BY DATE(created_at) ASC;

INSERT INTO fb_alert_test.summary_total_dailies (created_at, total_sub, total_unsub, total_renewal,total_revenue, total_active_sub, updated_at)
SELECT DATE(a.created_at) as created_at,  IFNULL(b.total, 0) + IFNULL(i.total, 0) + IFNULL(c.total, 0) as total_sub,
IFNULL(d.total, 0) as total_unsub, IFNULL(e.total, 0) + IFNULL(f.total, 0) as total_renewal, 
IFNULL(h.revenue, 0) + IFNULL(g.revenue, 0) as total_revenue, 
IFNULL(j.total_active_sub, 0) as total_active_sub, NOW()
FROM fb_alert_test.summary_revenues a
LEFT JOIN fb_alert_test.summary_revenues b ON b.created_at = a.created_at 
AND b.subject = 'FIRSTPUSH' AND b.status = 'SUCCESS'
LEFT JOIN fb_alert_test.summary_revenues i ON i.created_at = a.created_at 
AND i.subject = 'FIRSTPUSH' AND i.status = 'FAILED'
LEFT JOIN fb_alert_test.summary_revenues c ON c.created_at = a.created_at
AND c.subject = 'FREE' AND c.status = 'SUCCESS'
LEFT JOIN fb_alert_test.summary_revenues d ON d.created_at = a.created_at
AND d.subject = 'UNSUB'
LEFT JOIN fb_alert_test.summary_revenues e ON e.created_at = a.created_at
AND e.subject = 'RENEWAL' AND e.status = 'SUCCESS'
LEFT JOIN fb_alert_test.summary_revenues f ON f.created_at = a.created_at
AND f.subject = 'RENEWAL' AND f.status = 'FAILED'
LEFT JOIN fb_alert_test.summary_revenues g ON g.created_at = a.created_at
AND g.subject = 'FIRSTPUSH' AND g.status = 'SUCCESS'
LEFT JOIN fb_alert_test.summary_revenues h ON h.created_at = a.created_at
AND h.subject = 'RENEWAL' AND h.status = 'SUCCESS'
LEFT JOIN fb_alert_test.summary_dashboards j ON DATE(j.created_at) = DATE(a.created_at)
WHERE DATE(a.created_at) BETWEEN DATE('2024-12-01') AND DATE(NOW())
GROUP BY DATE(a.created_at)
ORDER BY DATE(a.created_at) ASC;
